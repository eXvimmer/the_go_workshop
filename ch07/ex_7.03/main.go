package main

import "fmt"

type record struct {
	key       string
	valueType string
	data      interface{}
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key

	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}

func main() {
	m := make(map[string]interface{})
	a := animal{name: "Pipa", category: "dog"}
	p := person{lastName: "Hayati", age: 29, isMarried: false}

	m["person"] = p
	m["animal"] = a
	m["age"] = 30
	m["isMarried"] = false
	m["lastName"] = "Hayati"
	m["pi"] = 3.14

	rs := []record{}

	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Println("key:", v.key)
		fmt.Println("data:", v.data)
		fmt.Println("type:", v.valueType)
		fmt.Println()
	}
}
