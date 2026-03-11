package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleStringToInt() {
	result := converter.StringToInt("123")
	fmt.Println(result)
	// Output: 123
}

func ExampleStringToInt_invalid() {
	result := converter.StringToInt("abc")
	fmt.Println(result) // алдаа гарвал 0 буцаана
	// Output: 0
}

func ExampleStringToUint() {
	result := converter.StringToUint("456")
	fmt.Println(result)
	// Output: 456
}

func ExampleStringToInt64() {
	result := converter.StringToInt64("9999999999")
	fmt.Println(result)
	// Output: 9999999999
}

func ExampleStringToFloat64() {
	result := converter.StringToFloat64("3.14")
	fmt.Println(result)
	// Output: 3.14
}

func ExampleStringToMap() {
	jsonStr := `{"name":"Бат","age":25}`
	m, err := converter.StringToMap(jsonStr)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(m["name"])
	// Output: Бат
}

func ExampleStringToMapArr() {
	jsonStr := `[{"id":1},{"id":2}]`
	arr, err := converter.StringToMapArr(jsonStr)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(len(arr))
	// Output: 2
}

func ExampleDateStringToTime() {
	t := converter.DateStringToTime("2024-01-15")
	fmt.Println(t.Year(), t.Month(), t.Day())
	// Output: 2024 January 15
}

func ExampleDateTimeStringToTime() {
	t := converter.DateTimeStringToTime("2024-01-15 10:30:00")
	fmt.Println(t.Hour(), t.Minute())
	// Output: 10 30
}

func ExampleStringToBool() {
	fmt.Println(converter.StringToBool("true"))
	fmt.Println(converter.StringToBool("false"))
	fmt.Println(converter.StringToBool("invalid"))
	// Output:
	// true
	// false
	// false
}
