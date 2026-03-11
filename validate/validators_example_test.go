package validate_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/validate"
)

func ExampleIsURL() {
	fmt.Println(validate.IsURL("https://example.com"))
	fmt.Println(validate.IsURL("not-a-url"))
	// Output:
	// true
	// false
}

func ExampleIsIP() {
	fmt.Println(validate.IsIP("192.168.1.1"))
	fmt.Println(validate.IsIP("::1"))
	fmt.Println(validate.IsIP("invalid"))
	// Output:
	// true
	// true
	// false
}

func ExampleIsStrongPassword() {
	fmt.Println(validate.IsStrongPassword("Passw0rd!"))
	fmt.Println(validate.IsStrongPassword("weak"))
	// Output:
	// true
	// false
}

func ExampleIsBankAccount() {
	fmt.Println(validate.IsBankAccount("1234567890"))
	fmt.Println(validate.IsBankAccount("123"))
	// Output:
	// true
	// false
}

func ExampleIsMongolianName() {
	fmt.Println(validate.IsMongolianName("Бат-Эрдэнэ"))
	fmt.Println(validate.IsMongolianName("John123"))
	// Output:
	// true
	// false
}
