package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	Message string
}

// NOTE: To be able to unmarshal into a struct, the struct field must be
// exportable. The struct's field name must be capitalized. Only fields that
// are exportable are visible externally, including the JSON unmarshaler. Only
// the exported fields will be in the JSON output; other fields are ignored.

func main() {
	data := []byte(`
    {
      "message": "Greeting fellow Gopher"
    }
  `)

	var v greeting

	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Message)
}
