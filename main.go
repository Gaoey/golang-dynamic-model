package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type MainStruct struct {
	X string `mapstructure:"field_x"`
	Y int    `mapstructure:"field_y"`
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
	var p map[string]interface{}
	v, _ := json.Marshal(input)
	_ = json.Unmarshal(v, &p)
	m.AddByGroup(p)
}

func (m *Model) Decode() interface{} {
	var s MainStruct
	_ = mapstructure.Decode(m.Payload, &s)
	fmt.Printf("payloadd = %#v, s result = %#v\n", m.Payload, s)
	return s
}

func main() {
	a1 := MainStruct{X: "xxx", Y: 15}

	test := New()
	test.AddByStruct(&a1)

	result := test.Decode()
	fmt.Printf("result = %#v", result)
}
