package types

import (
	"github.com/mitchellh/mapstructure"
)

type Map map[string]interface{}

func (m Map) Convert(o interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   o,
		TagName:  "input",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(m)
}

type WorkflowInput Map

func (i WorkflowInput) Convert(o interface{}) error {
	return Map(i).Convert(o)
}

type ActivityInput Map

func (i ActivityInput) Convert(o interface{}) error {
	return Map(i).Convert(o)
}
