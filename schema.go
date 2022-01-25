// Package jsonschema uses reflection to generate JSON Schemas from Go types [1].
//
// If json tags are present on struct fields, they will be used to infer
// property names and if a property is required (omitempty is present).
//
// [1] http://json-schema.org/latest/json-schema-validation.html
package jsonschema

import "encoding/json"

// Schema is the root schema.
// RFC draft-wright-json-schema-00, section 4.5
type Schema struct {
	*TypeDraft04
	Definitions Definitions
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(s.TypeDraft04)
	if err != nil {
		return nil, err
	}
	if s.Definitions == nil || len(s.Definitions) == 0 {
		return b, nil
	}
	d, err := json.Marshal(struct {
		Definitions Definitions `json:"definitions,omitempty"`
	}{s.Definitions})
	if err != nil {
		return nil, err
	}
	if len(b) == 2 {
		return d, nil
	} else {
		b[len(b)-1] = ','
		return append(b, d[1:]...), nil
	}
}
