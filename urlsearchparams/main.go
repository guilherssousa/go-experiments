package main

import (
	"fmt"
	"net/url"
)

type URLSearchParams map[string]any

func (p URLSearchParams) Has(key string) bool {
	_, ok := p[key]
	return ok
}

func (p URLSearchParams) Get(key string) any {
	return p[key]
}

func (p URLSearchParams) Set(key string, value any) {
	p[key] = value
}

func (p URLSearchParams) String() string {
	var str string
	for key, value := range p {
		value_type := fmt.Sprintf("%T", value)

		var value_str string
		switch value_type {
		case "string":
			value_str = value.(string)
		case "int":
			value_str = fmt.Sprintf("%d", value.(int))
		case "bool":
			value_str = fmt.Sprintf("%t", value.(bool))
		}

		str += fmt.Sprintf("%s=%s&", url.QueryEscape(key), url.QueryEscape(value_str))
	}
	return str[:len(str)-1]
}

func main() {
	params := URLSearchParams{
		"foo":         "bar",
		"baz":         123,
		"is he cool?": true,
	}

	params.Set("hello", "hello world")
	params.Set("foo", "bar2")

	// print params
	fmt.Println(params)
}
