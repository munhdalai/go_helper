package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleUintToString() {
	result := converter.UintToString(42)
	fmt.Println(result)
	// Output: 42
}

func ExampleIntToString() {
	result := converter.IntToString(-100)
	fmt.Println(result)
	// Output: -100
}

func ExampleInt64ToString() {
	result := converter.Int64ToString(9999999999)
	fmt.Println(result)
	// Output: 9999999999
}
