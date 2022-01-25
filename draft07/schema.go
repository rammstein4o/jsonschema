package draft07

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/orderedmap"
	"github.com/rammstein4o/jsonschema"
)

// Schema represents a JSON Schema object type.
type Schema struct {

	// Id see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.8.2
	Id string `json:"$id,omitempty"`

	// Comment see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.9
	Comment string `json:"$comment,omitempty"`

	// Schema see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.7
	Schema jsonschema.Version `json:"$schema,omitempty"`

	// Ref see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.8.3
	Ref string `json:"$ref,omitempty"`

	// Type see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.1.1
	Type string `json:"type,omitempty"`

	// Enum see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.1.2
	Enum []interface{} `json:"enum,omitempty"`

	// Const see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.1.3
	Const interface{} `json:"const,omitempty"`

	// MultipleOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.2.1
	MultipleOf uint64 `json:"multipleOf,omitempty"`

	// Maximum see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.2.2
	Maximum float64 `json:"maximum,omitempty"`

	// ExclusiveMaximum see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.2.3
	ExclusiveMaximum float64 `json:"exclusiveMaximum,omitempty"`

	// Minimum see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.2.4
	Minimum float64 `json:"minimum,omitempty"`

	// ExclusiveMinimum see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.2.5
	ExclusiveMinimum float64 `json:"exclusiveMinimum,omitempty"`

	// MaxLength see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.3.1
	MaxLength uint64 `json:"maxLength,omitempty"`

	// MinLength see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.3.2
	MinLength uint64 `json:"minLength,omitempty"`

	// Pattern see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.3.3
	Pattern string `json:"pattern,omitempty"`

	// Items see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.1
	Items *Schema `json:"items,omitempty"`

	// AdditionalItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.2
	AdditionalItems *Schema `json:"additionalItems,omitempty"`

	// MaxItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.3
	MaxItems uint64 `json:"maxItems,omitempty"`

	// MinItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.4
	MinItems uint64 `json:"minItems,omitempty"`

	// UniqueItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.5
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// Contains see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.6
	Contains *Schema `json:"contains,omitempty"`

	// MaxProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.1
	MaxProperties uint64 `json:"maxProperties,omitempty"`

	// MinProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.2
	MinProperties uint64 `json:"minProperties,omitempty"`

	// Required see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.3
	Required []string `json:"required,omitempty"`

	// Properties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.4
	Properties *orderedmap.OrderedMap `json:"properties,omitempty"`

	// PatternProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.5
	PatternProperties map[string]*Schema `json:"patternProperties,omitempty"`

	// AdditionalProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.6
	AdditionalProperties json.RawMessage `json:"additionalProperties,omitempty"`

	// Dependencies see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.7
	Dependencies map[string]*Schema `json:"dependencies,omitempty"`

	// PropertyNames see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.8
	PropertyNames *Schema `json:"propertyNames,omitempty"`

	// If see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.6.1
	If *Schema `json:"if,omitempty"`

	// Then see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.6.2
	Then *Schema `json:"then,omitempty"`

	// Else see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.6.3
	Else *Schema `json:"else,omitempty"`

	// AllOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.1
	AllOf []*Schema `json:"allOf,omitempty"`

	// AnyOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.2
	AnyOf []*Schema `json:"anyOf,omitempty"`

	// OneOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.3
	OneOf []*Schema `json:"oneOf,omitempty"`

	// Not see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.4
	Not *Schema `json:"not,omitempty"`

	// Format see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.7
	Format string `json:"format,omitempty"`

	// ContentEncoding see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.8.3
	ContentEncoding string `json:"contentEncoding,omitempty"`

	// ContentMediaType see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.8.4
	ContentMediaType string `json:"contentMediaType,omitempty"`

	// Definitions see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.9
	Definitions jsonschema.Definitions `json:"definitions,omitempty"`

	// Title see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.1
	Title string `json:"title,omitempty"`

	// Description see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.1
	Description string `json:"description,omitempty"`

	// Default see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.2
	Default interface{} `json:"default,omitempty"`

	// ReadOnly see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.3
	ReadOnly bool `json:"readOnly,omitempty"`

	// WriteOnly see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.3
	WriteOnly bool `json:"writeOnly,omitempty"`

	// Examples see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.10.4
	Examples []interface{} `json:"examples,omitempty"`

	Extras map[string]interface{} `json:"-"`
}

func (t *Schema) MarshalJSON() ([]byte, error) {
	type Schema_ Schema
	b, err := json.Marshal((*Schema_)(t))
	if err != nil {
		return nil, err
	}
	if t.Extras == nil || len(t.Extras) == 0 {
		return b, nil
	}
	m, err := json.Marshal(t.Extras)
	if err != nil {
		return nil, err
	}
	if len(b) == 2 {
		return m, nil
	} else {
		b[len(b)-1] = ','
		return append(b, m[1:]...), nil
	}
}

// read struct tags for array type keyworks
func (t *Schema) arrayKeywords(tags []string) {
	var defaultValues []interface{}
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "minItems":
				i, _ := strconv.ParseUint(val, 10, 64)
				t.MinItems = i
			case "maxItems":
				i, _ := strconv.ParseUint(val, 10, 64)
				t.MaxItems = i
			case "uniqueItems":
				t.UniqueItems = true
			case "default":
				defaultValues = append(defaultValues, val)
			case "enum":
				switch t.Items.Type {
				case "string":
					t.Items.Enum = append(t.Items.Enum, val)
				case "integer":
					i, _ := strconv.Atoi(val)
					t.Items.Enum = append(t.Items.Enum, i)
				case "number":
					f, _ := strconv.ParseFloat(val, 64)
					t.Items.Enum = append(t.Items.Enum, f)
				}
			}
		}
	}
	if len(defaultValues) > 0 {
		t.Default = defaultValues
	}
}

func (t *Schema) extraKeywords(tags []string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			t.setExtra(nameValue[0], nameValue[1])
		}
	}
}

func (t *Schema) setExtra(key, val string) {
	if t.Extras == nil {
		t.Extras = map[string]interface{}{}
	}
	if existingVal, ok := t.Extras[key]; ok {
		switch existingVal := existingVal.(type) {
		case string:
			t.Extras[key] = []string{existingVal, val}
		case []string:
			t.Extras[key] = append(existingVal, val)
		case int:
			t.Extras[key], _ = strconv.Atoi(val)
		}
	} else {
		switch key {
		case "minimum":
			t.Extras[key], _ = strconv.Atoi(val)
		default:
			t.Extras[key] = val
		}
	}
}

func (t *Schema) structKeywordsFromTags(f reflect.StructField, parentType *Schema, propertyName string) {
	t.Description = f.Tag.Get("jsonschema_description")
	tags := strings.Split(f.Tag.Get("jsonschema"), ",")
	t.genericKeywords(tags, parentType, propertyName)
	switch t.Type {
	case "string":
		t.stringKeywords(tags)
	case "number":
		t.numbericKeywords(tags)
	case "integer":
		t.numbericKeywords(tags)
	case "array":
		t.arrayKeywords(tags)
	}
	extras := strings.Split(f.Tag.Get("jsonschema_extras"), ",")
	t.extraKeywords(extras)
}

// read struct tags for generic keyworks
func (t *Schema) genericKeywords(tags []string, parentType *Schema, propertyName string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "title":
				t.Title = val
			case "description":
				t.Description = val
			case "type":
				t.Type = val
			case "oneof_required":
				var typeFound *Schema
				for i := range parentType.OneOf {
					if parentType.OneOf[i].Title == nameValue[1] {
						typeFound = parentType.OneOf[i]
					}
				}
				if typeFound == nil {
					typeFound = &Schema{
						Title:    nameValue[1],
						Required: []string{},
					}
					parentType.OneOf = append(parentType.OneOf, typeFound)
				}
				typeFound.Required = append(typeFound.Required, propertyName)
			case "oneof_type":
				if t.OneOf == nil {
					t.OneOf = make([]*Schema, 0, 1)
				}
				t.Type = ""
				types := strings.Split(nameValue[1], ";")
				for _, ty := range types {
					t.OneOf = append(t.OneOf, &Schema{
						Type: ty,
					})
				}
			case "enum":
				switch t.Type {
				case "string":
					t.Enum = append(t.Enum, val)
				case "integer":
					i, _ := strconv.Atoi(val)
					t.Enum = append(t.Enum, i)
				case "number":
					f, _ := strconv.ParseFloat(val, 64)
					t.Enum = append(t.Enum, f)
				}
			}
		}
	}
}

// read struct tags for string type keyworks
func (t *Schema) stringKeywords(tags []string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "minLength":
				i, _ := strconv.ParseUint(val, 10, 64)
				t.MinLength = i
			case "maxLength":
				i, _ := strconv.ParseUint(val, 10, 64)
				t.MaxLength = i
			case "pattern":
				t.Pattern = val
			case "format":
				switch val {
				case "date-time", "email", "hostname", "ipv4", "ipv6", "uri":
					t.Format = val
				}
			case "default":
				t.Default = val
			case "example":
				t.Examples = append(t.Examples, val)
			}
		}
	}
}

// read struct tags for numberic type keyworks
func (t *Schema) numbericKeywords(tags []string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "multipleOf":
				i, _ := strconv.ParseUint(val, 10, 64)
				t.MultipleOf = i
			case "minimum":
				i, _ := strconv.ParseFloat(val, 64)
				t.Minimum = i
			case "maximum":
				i, _ := strconv.ParseFloat(val, 64)
				t.Maximum = i
			case "exclusiveMaximum":
				b, _ := strconv.ParseFloat(val, 64)
				t.ExclusiveMaximum = b
			case "exclusiveMinimum":
				b, _ := strconv.ParseFloat(val, 64)
				t.ExclusiveMinimum = b
			case "default":
				i, _ := strconv.Atoi(val)
				t.Default = i
			case "example":
				if i, err := strconv.Atoi(val); err == nil {
					t.Examples = append(t.Examples, i)
				}
			}
		}
	}
}
