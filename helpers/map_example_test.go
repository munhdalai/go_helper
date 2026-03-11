package helpers_test

import (
	"fmt"
	"sort"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleMergeMaps() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}

	result := helpers.MergeMaps(m1, m2)
	fmt.Println(result["a"])
	fmt.Println(result["b"]) // m2-ийн утга давхцана
	fmt.Println(result["c"])
	// Output:
	// 1
	// 3
	// 4
}

func ExampleMapKeys() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := helpers.MapKeys(m)
	sort.Strings(keys)
	fmt.Println(keys)
	// Output: [a b c]
}

func ExampleMapValues() {
	m := map[string]int{"x": 10}
	values := helpers.MapValues(m)
	fmt.Println(values)
	// Output: [10]
}
