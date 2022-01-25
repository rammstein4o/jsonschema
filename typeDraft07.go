package jsonschema

import (
	"encoding/json"

	"github.com/iancoleman/orderedmap"
)

// TypeDraft07 represents a JSON Schema object type.
type TypeDraft07 struct {

	// Id see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.8.2
	Id string `json:"$id,omitempty"`

	// Comment see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.9
	Comment string `json:"$comment,omitempty"`

	// Schema see: https://json-schema.org/draft-07/json-schema-core.html#rfc.section.7
	Schema Version `json:"$schema,omitempty"`

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
	Items *TypeDraft07 `json:"items,omitempty"`

	// AdditionalItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.2
	AdditionalItems *TypeDraft07 `json:"additionalItems,omitempty"`

	// MaxItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.3
	MaxItems uint64 `json:"maxItems,omitempty"`

	// MinItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.4
	MinItems uint64 `json:"minItems,omitempty"`

	// UniqueItems see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.5
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// Contains see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.4.6
	Contains *TypeDraft07 `json:"contains,omitempty"`

	// MaxProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.1
	MaxProperties uint64 `json:"maxProperties,omitempty"`

	// MinProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.2
	MinProperties uint64 `json:"minProperties,omitempty"`

	// Required see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.3
	Required []string `json:"required,omitempty"`

	// Properties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.4
	Properties *orderedmap.OrderedMap `json:"properties,omitempty"`

	// PatternProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.5
	PatternProperties map[string]*TypeDraft07 `json:"patternProperties,omitempty"`

	// AdditionalProperties see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.6
	AdditionalProperties json.RawMessage `json:"additionalProperties,omitempty"`

	// Dependencies see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.7
	Dependencies map[string]*TypeDraft07 `json:"dependencies,omitempty"`

	// PropertyNames see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.5.8
	PropertyNames *TypeDraft07 `json:"propertyNames,omitempty"`

	// If see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.6.1
	If *TypeDraft07 `json:"if,omitempty"`

	// Then see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.6.2
	Then *TypeDraft07 `json:"then,omitempty"`

	// Else see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.6.3
	Else *TypeDraft07 `json:"else,omitempty"`

	// AllOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.1
	AllOf []*TypeDraft07 `json:"allOf,omitempty"`

	// AnyOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.2
	AnyOf []*TypeDraft07 `json:"anyOf,omitempty"`

	// OneOf see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.3
	OneOf []*TypeDraft07 `json:"oneOf,omitempty"`

	// Not see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.6.7.4
	Not *TypeDraft07 `json:"not,omitempty"`

	// Format see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.7
	Format string `json:"format,omitempty"`

	// ContentEncoding see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.8.3
	ContentEncoding string `json:"contentEncoding,omitempty"`

	// ContentMediaType see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.8.4
	ContentMediaType string `json:"contentMediaType,omitempty"`

	// Definitions see: https://json-schema.org/draft-07/json-schema-validation.html#rfc.section.9
	Definitions Definitions `json:"definitions,omitempty"`

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
