package jsonschema

import (
	"encoding/json"

	"github.com/iancoleman/orderedmap"
)

// TypeDraft202012 represents a JSON Schema object type.
type TypeDraft202012 struct {

	// Id see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.8.2.2
	Id string `json:"$id,omitempty"`

	// Comment see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.8.3
	Comment string `json:"$comment,omitempty"`

	// Schema see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.8.1.1
	Schema Version `json:"$schema,omitempty"`

	// Ref see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.8.2.3
	Ref string `json:"$ref,omitempty"`

	// Type see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.1.1
	Type string `json:"type,omitempty"`

	// Enum see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.1.2
	Enum []interface{} `json:"enum,omitempty"`

	// Const see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.1.3
	Const interface{} `json:"const,omitempty"`

	// MultipleOf see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.1
	MultipleOf uint64 `json:"multipleOf,omitempty"`

	// Maximum see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.2
	Maximum float64 `json:"maximum,omitempty"`

	// ExclusiveMaximum see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.3
	ExclusiveMaximum float64 `json:"exclusiveMaximum,omitempty"`

	// Minimum see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.4
	Minimum float64 `json:"minimum,omitempty"`

	// ExclusiveMinimum see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.2.5
	ExclusiveMinimum float64 `json:"exclusiveMinimum,omitempty"`

	// MaxLength see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.3.1
	MaxLength uint64 `json:"maxLength,omitempty"`

	// MinLength see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.3.2
	MinLength uint64 `json:"minLength,omitempty"`

	// Pattern see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.3.3
	Pattern string `json:"pattern,omitempty"`

	// PrefixItems see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.1
	PrefixItems *TypeDraft202012 `json:"prefixItems,omitempty"`

	// Items see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.1
	Items *TypeDraft202012 `json:"items,omitempty"`

	// AdditionalItems see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.1
	AdditionalItems *TypeDraft202012 `json:"additionalItems,omitempty"`

	// MaxItems see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.4.1
	MaxItems uint64 `json:"maxItems,omitempty"`

	// MinItems see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.4.2
	MinItems uint64 `json:"minItems,omitempty"`

	// UniqueItems see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.4.3
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// Contains see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.1
	Contains *TypeDraft202012 `json:"contains,omitempty"`

	// UniqueItems see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.4.4
	MaxContains uint64 `json:"maxContains,omitempty"` // section 6.4.6

	// UniqueItems see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.4.5
	MinContains uint64 `json:"minContains,omitempty"` // section 6.4.6

	// MaxProperties see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.5.1
	MaxProperties uint64 `json:"maxProperties,omitempty"`

	// MinProperties see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.5.2
	MinProperties uint64 `json:"minProperties,omitempty"`

	// Required see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.5.3
	Required []string `json:"required,omitempty"`

	// DependentRequired see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.5.4
	DependentRequired map[string][]string `json:"dependentRequired,omitempty"`

	// Properties see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.2
	Properties *orderedmap.OrderedMap `json:"properties,omitempty"`

	// PatternProperties see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.2
	PatternProperties map[string]*TypeDraft202012 `json:"patternProperties,omitempty"`

	// AdditionalProperties see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.2
	AdditionalProperties json.RawMessage `json:"additionalProperties,omitempty"`

	// Dependencies see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.6.5.7
	Dependencies map[string]*TypeDraft202012 `json:"dependencies,omitempty"`

	// PropertyNames see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.3.2
	PropertyNames *TypeDraft202012 `json:"propertyNames,omitempty"`

	// If see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.2
	If *TypeDraft202012 `json:"if,omitempty"`

	// Then see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.2
	Then *TypeDraft202012 `json:"then,omitempty"`

	// Else see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.2
	Else *TypeDraft202012 `json:"else,omitempty"`

	// Else see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.2
	DependentSchemas map[string]*TypeDraft202012 `json:"dependentSchemas,omitempty"`

	// AllOf see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.1
	AllOf []*TypeDraft202012 `json:"allOf,omitempty"`

	// AnyOf see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.1
	AnyOf []*TypeDraft202012 `json:"anyOf,omitempty"`

	// OneOf see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.1
	OneOf []*TypeDraft202012 `json:"oneOf,omitempty"`

	// Not see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.10.2.1
	Not *TypeDraft202012 `json:"not,omitempty"`

	// Format see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.7
	Format string `json:"format,omitempty"`

	// ContentEncoding see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.8.3
	ContentEncoding string `json:"contentEncoding,omitempty"`

	// ContentMediaType see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.8.4
	ContentMediaType string `json:"contentMediaType,omitempty"`

	// ContentSchema see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.8.4
	ContentSchema json.RawMessage `json:"contentSchema,omitempty"`

	// Definitions see: https://json-schema.org/draft/2020-12/json-schema-core.html#rfc.section.8.2.4
	Definitions Definitions `json:"$defs,omitempty"`

	// Title see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.1
	Title string `json:"title,omitempty"`

	// Description see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.1
	Description string `json:"description,omitempty"`

	// Default see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.2
	Default interface{} `json:"default,omitempty"`

	// Deprecated see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.3
	Deprecated bool `json:"deprecated,omitempty"`

	// ReadOnly see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.4
	ReadOnly bool `json:"readOnly,omitempty"`

	// WriteOnly see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.4
	WriteOnly bool `json:"writeOnly,omitempty"`

	// Examples see: https://json-schema.org/draft/2020-12/json-schema-validation.html#rfc.section.9.5
	Examples []interface{} `json:"examples,omitempty"`

	Extras map[string]interface{} `json:"-"`
}
