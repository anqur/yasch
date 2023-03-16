package types

import (
	"time"

	"github.com/alecthomas/units"
	"github.com/invopop/jsonschema"
)

type Duration struct{ Dur time.Duration }

func (Duration) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "string",
		Title:       "Duration",
		Description: "Go-compatible duration",
	}
}

func (d Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	dur, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.Dur = dur
	return nil
}

type Size struct{ Size units.Base2Bytes }

func (Size) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "string",
		Title:       "Size",
		Description: "Go-compatible data size",
	}
}

func (s Size) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var r string
	if err := unmarshal(&r); err != nil {
		return err
	}
	size, err := units.ParseBase2Bytes(r)
	if err != nil {
		return err
	}
	s.Size = size
	return nil
}
