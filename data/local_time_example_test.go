package data_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/munhdalai/go_helper/data"
)

func ExampleLocalTime_String() {
	t := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	lt := data.LocalTime(t)
	fmt.Println(lt.String())
	// Output: 2024-03-15 10:30:45
}

func ExampleLocalTime_ToTime() {
	t := time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC)
	lt := data.LocalTime(t)
	result := lt.ToTime()
	fmt.Println(result.Hour(), result.Minute())
	// Output: 10 30
}

func ExampleLocalTime_MarshalJSON() {
	type Log struct {
		Message   string         `json:"message"`
		CreatedAt data.LocalTime `json:"created_at"`
	}

	entry := Log{
		Message:   "Тест",
		CreatedAt: data.LocalTime(time.Date(2024, 6, 1, 14, 30, 0, 0, time.UTC)),
	}

	b, _ := json.Marshal(entry)
	fmt.Println(string(b))
	// Output: {"message":"Тест","created_at":"2024-06-01 14:30:00"}
}

func ExampleLocalTime_UnmarshalJSON() {
	type Log struct {
		Message   string         `json:"message"`
		CreatedAt data.LocalTime `json:"created_at"`
	}

	jsonStr := `{"message":"Тест","created_at":"2024-06-01 14:30:00"}`
	var entry Log
	json.Unmarshal([]byte(jsonStr), &entry)

	fmt.Println(entry.Message)
	fmt.Println(entry.CreatedAt.String())
	// Output:
	// Тест
	// 2024-06-01 14:30:00
}
