package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExamplePrettyJSON() {
	result, _ := converter.PrettyJSON(`{"name":"test","age":25}`)
	fmt.Println(result)
	// Output:
	// {
	//   "name": "test",
	//   "age": 25
	// }
}

func ExampleMinifyJSON() {
	result, _ := converter.MinifyJSON(`{  "name" : "test" ,  "age" : 25  }`)
	fmt.Println(result)
	// Output: {"name":"test","age":25}
}

func ExampleMergeJSON() {
	result, _ := converter.MergeJSON(`{"name":"test"}`, `{"age":25}`)
	fmt.Println(result)
	// Output: {"age":25,"name":"test"}
}
