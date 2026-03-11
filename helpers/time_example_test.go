package helpers_test

import (
	"fmt"
	"time"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleStartOfDay() {
	t := time.Date(2024, 6, 15, 14, 30, 0, 0, time.UTC)
	fmt.Println(helpers.StartOfDay(t))
	// Output: 2024-06-15 00:00:00 +0000 UTC
}

func ExampleEndOfDay() {
	t := time.Date(2024, 6, 15, 14, 30, 0, 0, time.UTC)
	fmt.Println(helpers.EndOfDay(t))
	// Output: 2024-06-15 23:59:59.999999999 +0000 UTC
}

func ExampleDaysBetween() {
	a := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	b := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println(helpers.DaysBetween(a, b))
	// Output: 9
}

func ExampleIsWeekend() {
	sat := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC) // Saturday
	mon := time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC) // Monday
	fmt.Println(helpers.IsWeekend(sat))
	fmt.Println(helpers.IsWeekend(mon))
	// Output:
	// true
	// false
}

func ExampleAddBusinessDays() {
	// 2024-06-14 is Friday
	t := time.Date(2024, 6, 14, 12, 0, 0, 0, time.UTC)
	result := helpers.AddBusinessDays(t, 3)
	fmt.Println(result.Format("2006-01-02"))
	// Output: 2024-06-19
}

func ExampleRelativeTime() {
	t := time.Now().Add(-2 * time.Hour)
	fmt.Println(helpers.RelativeTime(t))
	// Output: 2 цагийн өмнө
}

func ExampleIsToday() {
	fmt.Println(helpers.IsToday(time.Now()))
	// Output: true
}

func ExampleIsFuture() {
	fmt.Println(helpers.IsFuture(time.Now().Add(24 * time.Hour)))
	// Output: true
}

func ExampleIsPast() {
	fmt.Println(helpers.IsPast(time.Now().Add(-24 * time.Hour)))
	// Output: true
}
