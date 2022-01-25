package jsonschema

import (
	"encoding/json"
)

const (
	draft04Url     = "https://json-schema.org/draft-04/schema#"
	draft07Url     = "https://json-schema.org/draft-07/schema#"
	draft201909Url = "https://json-schema.org/draft/2019-09/schema#"
	draft202012Url = "https://json-schema.org/draft/2020-12/schema#"
)

type Version uint8

const (
	emptyDraft Version = iota
	DRAFT04
	DRAFT07
	DRAFT201909
	DRAFT202012
)

var versionToString = map[Version]string{
	DRAFT04:     draft04Url,
	DRAFT07:     draft07Url,
	DRAFT201909: draft201909Url,
	DRAFT202012: draft202012Url,
}

var versionToID = map[string]Version{
	draft04Url:     DRAFT04,
	draft07Url:     DRAFT07,
	draft201909Url: DRAFT201909,
	draft202012Url: DRAFT202012,
}

func (v Version) String() string {
	s, exist := versionToString[v]
	if !exist {
		return ""
	}
	return s
}

func (v Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *Version) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*v = DRAFT04
	if id, exist := versionToID[j]; exist {
		*v = id
	}
	return nil
}
