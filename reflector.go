package jsonschema

import (
	"reflect"
	"strings"

	"github.com/iancoleman/orderedmap"
)

// Reflect reflects to Schema from a value using the default Reflector
func Reflect(v interface{}) *Schema {
	return ReflectFromType(reflect.TypeOf(v))
}

// ReflectFromType generates root schema using the default Reflector
func ReflectFromType(t reflect.Type) *Schema {
	r := &Reflector{
		Version: DRAFT04,
	}
	return r.ReflectFromType(t)
}

// A Reflector reflects values into a Schema.
type Reflector struct {
	// Version defines target draft version
	Version Version

	// AllowAdditionalProperties will cause the Reflector to generate a schema
	// with additionalProperties to 'true' for all struct types. This means
	// the presence of additional keys in JSON objects will not cause validation
	// to fail. Note said additional keys will simply be dropped when the
	// validated JSON is unmarshaled.
	AllowAdditionalProperties bool

	// RequiredFromJSONSchemaTags will cause the Reflector to generate a schema
	// that requires any key tagged with `jsonschema:required`, overriding the
	// default of requiring any key *not* tagged with `json:,omitempty`.
	RequiredFromJSONSchemaTags bool

	// YAMLEmbeddedStructs will cause the Reflector to generate a schema that does
	// not inline embedded structs. This should be enabled if the JSON schemas are
	// used with yaml.Marshal/Unmarshal.
	YAMLEmbeddedStructs bool

	// Prefer yaml: tags over json: tags to generate the schema even if json: tags
	// are present
	PreferYAMLSchema bool

	// ExpandedStruct will cause the toplevel definitions of the schema not
	// be referenced itself to a definition.
	ExpandedStruct bool

	// Do not reference definitions.
	// All types are still registered under the "definitions" top-level object,
	// but instead of $ref fields in containing types, the entire definition
	// of the contained type is inserted.
	// This will cause the entire structure of types to be output in one tree.
	DoNotReference bool

	// Use package paths as well as type names, to avoid conflicts.
	// Without this setting, if two packages contain a type with the same name,
	// and both are present in a schema, they will conflict and overwrite in
	// the definition map and produce bad output.  This is particularly
	// noticeable when using DoNotReference.
	FullyQualifyTypeNames bool

	// IgnoredTypes defines a slice of types that should be ignored in the schema,
	// switching to just allowing additional properties instead.
	IgnoredTypes []interface{}

	// TypeMapper is a function that can be used to map custom Go types to jsonschema types.
	TypeMapper func(reflect.Type) *TypeDraft04

	// TypeNamer allows customizing of type names
	TypeNamer func(reflect.Type) string

	// AdditionalFields allows adding structfields for a given type
	AdditionalFields func(reflect.Type) []reflect.StructField
}

// Reflect reflects to Schema from a value.
func (r *Reflector) Reflect(v interface{}) *Schema {
	return r.ReflectFromType(reflect.TypeOf(v))
}

// ReflectFromType generates root schema
func (r *Reflector) ReflectFromType(t reflect.Type) *Schema {
	definitions := Definitions{}
	if r.ExpandedStruct {
		st := &TypeDraft04{
			Version:              r.Version,
			Type:                 "object",
			Properties:           orderedmap.New(),
			AdditionalProperties: []byte("false"),
		}
		if r.AllowAdditionalProperties {
			st.AdditionalProperties = []byte("true")
		}
		r.reflectStructFields(st, definitions, t)
		r.reflectStruct(definitions, t)
		delete(definitions, r.typeName(t))
		return &Schema{TypeDraft04: st, Definitions: definitions}
	}

	s := &Schema{
		TypeDraft04: r.reflectTypeToSchema(definitions, t),
		Definitions: definitions,
	}
	return s
}

func (r *Reflector) reflectTypeToSchema(definitions Definitions, t reflect.Type) *TypeDraft04 {
	// Already added to definitions?
	if _, ok := definitions[r.typeName(t)]; ok && !r.DoNotReference {
		return &TypeDraft04{Ref: "#/definitions/" + r.typeName(t)}
	}

	if r.TypeMapper != nil {
		if t := r.TypeMapper(t); t != nil {
			return t
		}
	}

	if rt := r.reflectCustomType(definitions, t); rt != nil {
		return rt
	}

	// jsonpb will marshal protobuf enum options as either strings or integers.
	// It will unmarshal either.
	if t.Implements(protoEnumType) {
		return &TypeDraft04{OneOf: []*TypeDraft04{
			{Type: "string"},
			{Type: "integer"},
		}}
	}

	// Defined format types for JSON Schema Validation
	// RFC draft-wright-json-schema-validation-00, section 7.3
	// TODO email RFC section 7.3.2, hostname RFC section 7.3.3, uriref RFC section 7.3.7
	if t == ipType {
		// TODO differentiate ipv4 and ipv6 RFC section 7.3.4, 7.3.5
		return &TypeDraft04{Type: "string", Format: "ipv4"} // ipv4 RFC section 7.3.4
	}

	switch t.Kind() {
	case reflect.Struct:
		switch t {
		case timeType: // date-time RFC section 7.3.1
			return &TypeDraft04{Type: "string", Format: "date-time"}
		case uriType: // uri RFC section 7.3.6
			return &TypeDraft04{Type: "string", Format: "uri"}
		default:
			return r.reflectStruct(definitions, t)
		}

	case reflect.Map:
		switch t.Key().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			rt := &TypeDraft04{
				Type: "object",
				PatternProperties: map[string]*TypeDraft04{
					"^[0-9]+$": r.reflectTypeToSchema(definitions, t.Elem()),
				},
				AdditionalProperties: []byte("false"),
			}
			return rt
		}

		rt := &TypeDraft04{
			Type: "object",
			PatternProperties: map[string]*TypeDraft04{
				".*": r.reflectTypeToSchema(definitions, t.Elem()),
			},
		}
		delete(rt.PatternProperties, "additionalProperties")
		return rt

	case reflect.Slice, reflect.Array:
		returnType := &TypeDraft04{}
		if t == rawMessageType {
			return &TypeDraft04{
				AdditionalProperties: []byte("true"),
			}
		}
		if t.Kind() == reflect.Array {
			returnType.MinItems = t.Len()
			returnType.MaxItems = returnType.MinItems
		}
		if t.Kind() == reflect.Slice && t.Elem() == byteSliceType.Elem() {
			returnType.Type = "string"
			returnType.Media = &TypeDraft04{BinaryEncoding: "base64"}
			return returnType
		}
		returnType.Type = "array"
		returnType.Items = r.reflectTypeToSchema(definitions, t.Elem())
		return returnType

	case reflect.Interface:
		return &TypeDraft04{
			AdditionalProperties: []byte("true"),
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &TypeDraft04{Type: "integer"}

	case reflect.Float32, reflect.Float64:
		return &TypeDraft04{Type: "number"}

	case reflect.Bool:
		return &TypeDraft04{Type: "boolean"}

	case reflect.String:
		return &TypeDraft04{Type: "string"}

	case reflect.Ptr:
		return r.reflectTypeToSchema(definitions, t.Elem())
	}
	panic("unsupported type " + t.String())
}

func (r *Reflector) reflectCustomType(definitions Definitions, t reflect.Type) *TypeDraft04 {
	if t.Kind() == reflect.Ptr {
		return r.reflectCustomType(definitions, t.Elem())
	}

	if t.Implements(customType) {
		v := reflect.New(t)
		o := v.Interface().(customSchemaType)
		st := o.JSONSchemaType()
		definitions[r.typeName(t)] = st
		if r.DoNotReference {
			return st
		} else {
			return &TypeDraft04{
				Version: r.Version,
				Ref:     "#/definitions/" + r.typeName(t),
			}
		}
	}

	return nil
}

// Reflects a struct to a JSON Schema type.
func (r *Reflector) reflectStruct(definitions Definitions, t reflect.Type) *TypeDraft04 {
	if st := r.reflectCustomType(definitions, t); st != nil {
		return st
	}

	for _, ignored := range r.IgnoredTypes {
		if reflect.TypeOf(ignored) == t {
			st := &TypeDraft04{
				Type:                 "object",
				Properties:           orderedmap.New(),
				AdditionalProperties: []byte("true"),
			}
			definitions[r.typeName(t)] = st

			if r.DoNotReference {
				return st
			} else {
				return &TypeDraft04{
					Version: r.Version,
					Ref:     "#/definitions/" + r.typeName(t),
				}
			}
		}
	}

	st := &TypeDraft04{
		Type:                 "object",
		Properties:           orderedmap.New(),
		AdditionalProperties: []byte("false"),
	}
	if r.AllowAdditionalProperties {
		st.AdditionalProperties = []byte("true")
	}
	definitions[r.typeName(t)] = st
	r.reflectStructFields(st, definitions, t)

	if r.DoNotReference {
		return st
	} else {
		return &TypeDraft04{
			Version: r.Version,
			Ref:     "#/definitions/" + r.typeName(t),
		}
	}
}

func (r *Reflector) reflectStructFields(st *TypeDraft04, definitions Definitions, t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}

	var getFieldDocString customGetFieldDocString
	if t.Implements(customStructGetFieldDocString) {
		v := reflect.New(t)
		o := v.Interface().(customSchemaGetFieldDocString)
		getFieldDocString = o.GetFieldDocString
	}

	handleField := func(f reflect.StructField) {
		name, shouldEmbed, required, nullable := r.reflectFieldName(f)
		// if anonymous and exported type should be processed recursively
		// current type should inherit properties of anonymous one
		if name == "" {
			if shouldEmbed {
				r.reflectStructFields(st, definitions, f.Type)
			}
			return
		}

		property := r.reflectTypeToSchema(definitions, f.Type)
		property.structKeywordsFromTags(f, st, name)
		if getFieldDocString != nil {
			property.Description = getFieldDocString(f.Name)
		}

		if nullable {
			property = &TypeDraft04{
				OneOf: []*TypeDraft04{
					property,
					{
						Type: "null",
					},
				},
			}
		}

		st.Properties.Set(name, property)
		if required {
			st.Required = append(st.Required, name)
		}
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		handleField(f)
	}
	if r.AdditionalFields != nil {
		if af := r.AdditionalFields(t); af != nil {
			for _, sf := range af {
				handleField(sf)
			}
		}
	}
}

func (r *Reflector) reflectFieldName(f reflect.StructField) (string, bool, bool, bool) {
	jsonTags, exist := f.Tag.Lookup("json")
	yamlTags, yamlExist := f.Tag.Lookup("yaml")
	if !exist || r.PreferYAMLSchema {
		jsonTags = yamlTags
		exist = yamlExist
	}

	jsonTagsList := strings.Split(jsonTags, ",")
	yamlTagsList := strings.Split(yamlTags, ",")

	if ignoredByJSONTags(jsonTagsList) {
		return "", false, false, false
	}

	jsonSchemaTags := strings.Split(f.Tag.Get("jsonschema"), ",")
	if ignoredByJSONSchemaTags(jsonSchemaTags) {
		return "", false, false, false
	}

	name := f.Name
	required := requiredFromJSONTags(jsonTagsList)

	if r.RequiredFromJSONSchemaTags {
		required = requiredFromJSONSchemaTags(jsonSchemaTags)
	}

	nullable := nullableFromJSONSchemaTags(jsonSchemaTags)

	if jsonTagsList[0] != "" {
		name = jsonTagsList[0]
	}

	// field not anonymous and not export has no export name
	if !f.Anonymous && f.PkgPath != "" {
		name = ""
	}

	embed := false

	// field anonymous but without json tag should be inherited by current type
	if f.Anonymous && !exist {
		if !r.YAMLEmbeddedStructs {
			name = ""
			embed = true
		} else {
			name = strings.ToLower(name)
		}
	}

	if yamlExist && inlineYAMLTags(yamlTagsList) {
		name = ""
		embed = true
	}

	return name, embed, required, nullable
}

func (r *Reflector) typeName(t reflect.Type) string {
	if r.TypeNamer != nil {
		if name := r.TypeNamer(t); name != "" {
			return name
		}
	}
	if r.FullyQualifyTypeNames {
		return t.PkgPath() + "." + t.Name()
	}
	return t.Name()
}
