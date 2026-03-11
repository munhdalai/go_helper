package helpers_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleMin() {
	fmt.Println(helpers.Min(3, 7))
	fmt.Println(helpers.Min(10.5, 2.3))
	// Output:
	// 3
	// 2.3
}

func ExampleMax() {
	fmt.Println(helpers.Max(3, 7))
	fmt.Println(helpers.Max(10.5, 2.3))
	// Output:
	// 7
	// 10.5
}

func ExampleClamp() {
	fmt.Println(helpers.Clamp(5, 1, 10))
	fmt.Println(helpers.Clamp(-5, 0, 100))
	fmt.Println(helpers.Clamp(150, 0, 100))
	// Output:
	// 5
	// 0
	// 100
}

func ExampleSum() {
	fmt.Println(helpers.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(helpers.Sum([]float64{1.5, 2.5, 3.0}))
	// Output:
	// 15
	// 7
}

func ExampleAverage() {
	fmt.Println(helpers.Average([]int{1, 2, 3, 4, 5}))
	fmt.Println(helpers.Average([]float64{10.0, 20.0, 30.0}))
	// Output:
	// 3
	// 20
}

func ExampleAbs() {
	fmt.Println(helpers.Abs(-5))
	fmt.Println(helpers.Abs(3.14))
	fmt.Println(helpers.Abs(-2.5))
	// Output:
	// 5
	// 3.14
	// 2.5
}
