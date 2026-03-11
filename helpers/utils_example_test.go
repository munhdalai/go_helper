package helpers_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleGeneratePassword() {
	hash := helpers.GeneratePassword("password")
	fmt.Println(len(hash)) // SHA1 hex нь 40 тэмдэгт
	// Output: 40
}

func ExampleToSnakeCase() {
	fmt.Println(helpers.ToSnakeCase("CamelCase"))
	fmt.Println(helpers.ToSnakeCase("myVariableName"))
	fmt.Println(helpers.ToSnakeCase("HTTPResponse"))
	// Output:
	// camel_case
	// my_variable_name
	// http_response
}

func ExampleRandString() {
	s := helpers.RandString(10)
	fmt.Println(len(s))
	// Output: 10
}

func ExampleStringInSlice() {
	list := []string{"apple", "banana", "cherry"}
	fmt.Println(helpers.StringInSlice("banana", list))
	fmt.Println(helpers.StringInSlice("grape", list))
	// Output:
	// true
	// false
}

func ExampleGenerateOTP() {
	otp := helpers.GenerateOTP(6)
	// 6 оронтой тоо (100000-999999)
	fmt.Println(otp >= 100000 && otp <= 999999)
	// Output: true
}

func ExampleGenerateMD5Hash() {
	hash := helpers.GenerateMD5Hash("hello", "world")
	fmt.Println(hash)
	// Output: fc5e038d38a57032085441e7fe7010b0
}

func ExampleUniqueIntSlice() {
	slice := []int{1, 2, 2, 3, 3, 3}
	result := helpers.UniqueIntSlice(slice)
	fmt.Println(result)
	// Output: [1 2 3]
}

func ExampleUniqueStringSlice() {
	slice := []string{"a", "b", "a", "c", "b"}
	result := helpers.UniqueStringSlice(slice)
	fmt.Println(result)
	// Output: [a b c]
}

func ExampleRoundFloat() {
	fmt.Println(helpers.RoundFloat(3.14159))
	fmt.Println(helpers.RoundFloat(2.999))
	// Output:
	// 3.14
	// 2.99
}

func ExampleMask() {
	// maskType=1: сондгой индекстэй тэмдэгтүүдийг * болгоно
	fmt.Println(helpers.Mask("12345678", 1))
	fmt.Println(helpers.Mask("ABCDEF", 1))
	// Output:
	// 1*3*5*7*
	// A*C*E*
}
