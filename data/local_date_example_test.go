package data_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/munhdalai/go_helper/data"
)

func ExampleLocalDate_String() {
	t := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	d := data.LocalDate(t)
	fmt.Println(d.String())
	// Output: 2024-03-15
}

func ExampleLocalDate_ToTime() {
	t := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	d := data.LocalDate(t)
	result := d.ToTime()
	fmt.Println(result.Year(), result.Month(), result.Day())
	// Output: 2024 March 15
}

func ExampleLocalDate_MarshalJSON() {
	type Event struct {
		Name string         `json:"name"`
		Date data.LocalDate `json:"date"`
	}

	event := Event{
		Name: "Тест",
		Date: data.LocalDate(time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)),
	}

	b, _ := json.Marshal(event)
	fmt.Println(string(b))
	// Output: {"name":"Тест","date":"2024-06-01"}
}

func ExampleLocalDate_UnmarshalJSON() {
	type Event struct {
		Name string         `json:"name"`
		Date data.LocalDate `json:"date"`
	}

	jsonStr := `{"name":"Тест","date":"2024-06-01"}`
	var event Event
	json.Unmarshal([]byte(jsonStr), &event)

	fmt.Println(event.Name)
	fmt.Println(event.Date.String())
	// Output:
	// Тест
	// 2024-06-01
}
