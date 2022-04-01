package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type greeting struct {
	SomeMessage string `json:"message"`
}

// NOTE: To be able to unmarshal into a struct, the struct field must be
// exportable. The struct's field name must be capitalized. Only fields that
// are exportable are visible externally, including the JSON unmarshaler. Only
// the exported fields will be in the JSON output; other fields are ignored.

// NOTE: Using struct tags gives us more control. We can now name our struct
// field name anything as long as it is exportable.

func main() {
	data := []byte(`
    {
      "message": "Greeting fellow Gopher"
    }
  `)

	var v greeting

	if !json.Valid(data) {
		fmt.Printf("JSON is not valid: %s\n", data)
		os.Exit(1)
	}

	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.SomeMessage)
}
