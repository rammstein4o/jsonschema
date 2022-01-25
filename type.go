package jsonschema

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/orderedmap"
)

// Type represents a JSON Schema object type.
type Type struct {
	// RFC draft-wright-json-schema-00
	Version Version `json:"$schema,omitempty"` // section 6.1
	Ref     string  `json:"$ref,omitempty"`    // section 7
	// RFC draft-wright-json-schema-validation-00, section 5
	MultipleOf           int                    `json:"multipleOf,omitempty"`           // section 5.1
	Maximum              int                    `json:"maximum,omitempty"`              // section 5.2
	ExclusiveMaximum     bool                   `json:"exclusiveMaximum,omitempty"`     // section 5.3
	Minimum              int                    `json:"minimum,omitempty"`              // section 5.4
	ExclusiveMinimum     bool                   `json:"exclusiveMinimum,omitempty"`     // section 5.5
	MaxLength            int                    `json:"maxLength,omitempty"`            // section 5.6
	MinLength            int                    `json:"minLength,omitempty"`            // section 5.7
	Pattern              string                 `json:"pattern,omitempty"`              // section 5.8
	AdditionalItems      *Type                  `json:"additionalItems,omitempty"`      // section 5.9
	Items                *Type                  `json:"items,omitempty"`                // section 5.9
	MaxItems             int                    `json:"maxItems,omitempty"`             // section 5.10
	MinItems             int                    `json:"minItems,omitempty"`             // section 5.11
	UniqueItems          bool                   `json:"uniqueItems,omitempty"`          // section 5.12
	MaxProperties        int                    `json:"maxProperties,omitempty"`        // section 5.13
	MinProperties        int                    `json:"minProperties,omitempty"`        // section 5.14
	Required             []string               `json:"required,omitempty"`             // section 5.15
	Properties           *orderedmap.OrderedMap `json:"properties,omitempty"`           // section 5.16
	PatternProperties    map[string]*Type       `json:"patternProperties,omitempty"`    // section 5.17
	AdditionalProperties json.RawMessage        `json:"additionalProperties,omitempty"` // section 5.18
	Dependencies         map[string]*Type       `json:"dependencies,omitempty"`         // section 5.19
	Enum                 []interface{}          `json:"enum,omitempty"`                 // section 5.20
	Type                 string                 `json:"type,omitempty"`                 // section 5.21
	AllOf                []*Type                `json:"allOf,omitempty"`                // section 5.22
	AnyOf                []*Type                `json:"anyOf,omitempty"`                // section 5.23
	OneOf                []*Type                `json:"oneOf,omitempty"`                // section 5.24
	Not                  *Type                  `json:"not,omitempty"`                  // section 5.25
	Definitions          Definitions            `json:"definitions,omitempty"`          // section 5.26
	// RFC draft-wright-json-schema-validation-00, section 6, 7
	Title       string        `json:"title,omitempty"`       // section 6.1
	Description string        `json:"description,omitempty"` // section 6.1
	Default     interface{}   `json:"default,omitempty"`     // section 6.2
	Format      string        `json:"format,omitempty"`      // section 7
	Examples    []interface{} `json:"examples,omitempty"`    // section 7.4
	// RFC draft-wright-json-schema-hyperschema-00, section 4
	Media          *Type  `json:"media,omitempty"`          // section 4.3
	BinaryEncoding string `json:"binaryEncoding,omitempty"` // section 4.3

	Extras map[string]interface{} `json:"-"`
}

func (t *Type) MarshalJSON() ([]byte, error) {
	type Type_ Type
	b, err := json.Marshal((*Type_)(t))
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
func (t *Type) arrayKeywords(tags []string) {
	var defaultValues []interface{}
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "minItems":
				i, _ := strconv.Atoi(val)
				t.MinItems = i
			case "maxItems":
				i, _ := strconv.Atoi(val)
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

func (t *Type) extraKeywords(tags []string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			t.setExtra(nameValue[0], nameValue[1])
		}
	}
}

func (t *Type) setExtra(key, val string) {
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

func (t *Type) structKeywordsFromTags(f reflect.StructField, parentType *Type, propertyName string) {
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
func (t *Type) genericKeywords(tags []string, parentType *Type, propertyName string) {
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
				var typeFound *Type
				for i := range parentType.OneOf {
					if parentType.OneOf[i].Title == nameValue[1] {
						typeFound = parentType.OneOf[i]
					}
				}
				if typeFound == nil {
					typeFound = &Type{
						Title:    nameValue[1],
						Required: []string{},
					}
					parentType.OneOf = append(parentType.OneOf, typeFound)
				}
				typeFound.Required = append(typeFound.Required, propertyName)
			case "oneof_type":
				if t.OneOf == nil {
					t.OneOf = make([]*Type, 0, 1)
				}
				t.Type = ""
				types := strings.Split(nameValue[1], ";")
				for _, ty := range types {
					t.OneOf = append(t.OneOf, &Type{
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
func (t *Type) stringKeywords(tags []string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "minLength":
				i, _ := strconv.Atoi(val)
				t.MinLength = i
			case "maxLength":
				i, _ := strconv.Atoi(val)
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
func (t *Type) numbericKeywords(tags []string) {
	for _, tag := range tags {
		nameValue := strings.Split(tag, "=")
		if len(nameValue) == 2 {
			name, val := nameValue[0], nameValue[1]
			switch name {
			case "multipleOf":
				i, _ := strconv.Atoi(val)
				t.MultipleOf = i
			case "minimum":
				i, _ := strconv.Atoi(val)
				t.Minimum = i
			case "maximum":
				i, _ := strconv.Atoi(val)
				t.Maximum = i
			case "exclusiveMaximum":
				b, _ := strconv.ParseBool(val)
				t.ExclusiveMaximum = b
			case "exclusiveMinimum":
				b, _ := strconv.ParseBool(val)
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
