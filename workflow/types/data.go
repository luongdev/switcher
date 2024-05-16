package types

import (
	"encoding/json"
	"fmt"
	"github.com/luongdev/switcher/workflow/enums"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strings"
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

func (m Map) Bytes() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (m Map) Set(s interface{}) error {
	typ := reflect.TypeOf(s)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}

	val := reflect.ValueOf(s).Elem()
	for i := 0; i < typ.Elem().NumField(); i++ {
		field := typ.Elem().Field(i)
		tag := strings.Split(field.Tag.Get("json"), ",")[0]
		if tag == "" {
			tag = field.Name
		}

		if _, ok := m[tag]; !ok {
			m[tag] = val.Field(i).Interface()
		}
	}

	return nil
}

type Metadata map[enums.Field]interface{}
