package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type MainStruct struct {
	X string `mapstructure:"field_x" json:"field_x"`
	Y int    `mapstructure:"field_y" json:"field_y"`
	Z string `mapstructure:"field_z" json:"field_z"`
}

type Model struct {
	Payload map[string]interface{}
}

func New() *Model {
	return &Model{
		Payload: make(map[string]interface{}),
	}
}

func (m *Model) Add(key string, value interface{}) {
	m.Payload[key] = value
}

func (m *Model) Remove(key string) {
	delete(m.Payload, key)
}

func (m *Model) AddByGroup(p map[string]interface{}) {
	for k, v := range p {
		m.Add(k, v)
	}
}

func (m *Model) AddByStruct(input interface{}) {
	v := reflect.ValueOf(input)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := string(field.Tag.Get("mapstructure"))
		m.Add(tag, v.Field(i).Interface())
	}
}

// Decode to main struct
func (m *Model) DecodeAs(s interface{}) interface{} {
	_ = mapstructure.Decode(m.Payload, &s)
	return s
}

func main() {
	a1 := MainStruct{X: "xxx", Y: 15}

	test := New()
	test.AddByStruct(a1)
	test.Add("field_z", "zzz")

	fmt.Printf("%v\n", test.Payload)

	result := test.DecodeAs(MainStruct{})

	b, _ := json.Marshal(result)
	fmt.Printf("%s", string(b))
}
