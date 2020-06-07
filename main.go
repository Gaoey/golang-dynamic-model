package main

import "fmt"

type MainStruct struct {
	x string `mapstructure:"field_x"`
	y int `mapstructure:"field_y"`
}

type Model struct{
	payload map[string]interface{}
}

func New() *Model{
	return &Model
}

func (m *Model) Add(key string, value interface{}) {
	m.payload[key] = value
}

func (m *Model) Remove(key string) {
	delete(m.payload, key)
}

func (m *Model) AddByGroup(p map[string]interface{}) {
	for k, v := range p {
		m.Add(k,v)
	}
}

func (m *Model) AddByStruct(input interface{}) {
	var p map[string]interface{}
	json.Marshal(&input, &p)
	m.AddByGroup(p)
}

func (m *Model) Decode() interface{}{
	map
}

func main()  {
	a1 := struct {
		x string
		y int
	}{
			x :"xxx",
			y: 0,
	}

	test := 
}