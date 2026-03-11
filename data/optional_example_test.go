package data_test

import (
	"encoding/json"
	"fmt"

	"github.com/munhdalai/go_helper/data"
)

func ExampleNewNull() {
	n := data.NewNull(42)
	fmt.Println(n.Val, n.Valid)
	// Output: 42 true
}

func ExampleNullFrom() {
	val := "hello"
	n := data.NullFrom(&val)
	fmt.Println(n.Val, n.Valid)
	n2 := data.NullFrom[string](nil)
	fmt.Println(n2.Valid)
	// Output:
	// hello true
	// false
}

func ExampleNull_IsNull() {
	n := data.NewNull(42)
	fmt.Println(n.IsNull())
	var empty data.Null[int]
	fmt.Println(empty.IsNull())
	// Output:
	// false
	// true
}

func ExampleNull_ValueOrDefault() {
	n := data.NewNull(42)
	fmt.Println(n.ValueOrDefault(0))
	var empty data.Null[int]
	fmt.Println(empty.ValueOrDefault(99))
	// Output:
	// 42
	// 99
}

func ExampleNull_Ptr() {
	n := data.NewNull(42)
	fmt.Println(*n.Ptr())
	var empty data.Null[int]
	fmt.Println(empty.Ptr())
	// Output:
	// 42
	// <nil>
}

func ExampleNull_MarshalJSON() {
	n := data.NewNull("hello")
	b, _ := json.Marshal(n)
	fmt.Println(string(b))
	var empty data.Null[string]
	b2, _ := json.Marshal(empty)
	fmt.Println(string(b2))
	// Output:
	// "hello"
	// null
}

func ExampleNull_UnmarshalJSON() {
	var n data.Null[int]
	json.Unmarshal([]byte("42"), &n)
	fmt.Println(n.Val, n.Valid)
	var n2 data.Null[int]
	json.Unmarshal([]byte("null"), &n2)
	fmt.Println(n2.Valid)
	// Output:
	// 42 true
	// false
}
