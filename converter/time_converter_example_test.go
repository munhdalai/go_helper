package converter_test

import (
	"fmt"
	"time"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleTimeToDateString() {
	t := time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC)
	result := converter.TimeToDateString(t)
	fmt.Println(result)
	// Output: 2024-03-15
}

func ExampleTimeToDateTimeString() {
	t := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	result := converter.TimeToDateTimeString(t)
	fmt.Println(result)
	// Output: 2024-03-15 10:30:45
}

func ExampleTimeToTimeString() {
	t := time.Date(2024, 3, 15, 10, 30, 45, 0, time.UTC)
	result := converter.TimeToTimeString(t)
	fmt.Println(result)
	// Output: 10:30:45
}
