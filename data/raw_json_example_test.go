package data_test

import (
	"encoding/json"
	"fmt"

	"github.com/munhdalai/go_helper/data"
)

func ExampleJSONRaw_MarshalJSON() {
	type Record struct {
		ID       int          `json:"id"`
		Metadata data.JSONRaw `json:"metadata"`
	}

	record := Record{
		ID:       1,
		Metadata: data.JSONRaw(`{"key":"value","count":42}`),
	}

	b, _ := json.Marshal(record)
	fmt.Println(string(b))
	// Output: {"id":1,"metadata":{"key":"value","count":42}}
}

func ExampleJSONRaw_UnmarshalJSON() {
	type Record struct {
		ID       int          `json:"id"`
		Metadata data.JSONRaw `json:"metadata"`
	}

	jsonStr := `{"id":1,"metadata":{"key":"value"}}`
	var record Record
	json.Unmarshal([]byte(jsonStr), &record)

	fmt.Println(record.ID)
	fmt.Println(string(record.Metadata))
	// Output:
	// 1
	// {"key":"value"}
}
