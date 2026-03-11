package helpers_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleRandInt() {
	result := helpers.RandInt(1, 10)
	fmt.Println(result >= 1 && result <= 10)
	// Output: true
}

func ExampleRandFloat() {
	result := helpers.RandFloat(1.0, 10.0)
	fmt.Println(result >= 1.0 && result < 10.0)
	// Output: true
}

func ExampleRandElement() {
	items := []string{"a", "b", "c"}
	result := helpers.RandElement(items)
	fmt.Println(result == "a" || result == "b" || result == "c")
	// Output: true
}

func ExampleShuffle() {
	items := []int{1, 2, 3, 4, 5}
	helpers.Shuffle(items)
	fmt.Println(len(items))
	// Output: 5
}
