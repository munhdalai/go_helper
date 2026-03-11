package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleToPointer() {
	val := 42
	ptr := converter.ToPointer(val)
	fmt.Println(*ptr)
	// Output: 42
}

func ExampleToPointer_string() {
	name := "Бат"
	ptr := converter.ToPointer(name)
	fmt.Println(*ptr)
	// Output: Бат
}

func ExampleFromPointer() {
	val := 42
	ptr := &val
	result := converter.FromPointer(ptr)
	fmt.Println(result)
	// Output: 42
}

func ExampleFromPointer_nil() {
	var ptr *int
	result := converter.FromPointer(ptr)
	fmt.Println(result) // nil бол zero value буцаана
	// Output: 0
}
