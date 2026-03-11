package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleFormatPhone() {
	fmt.Println(converter.FormatPhone("99112233"))
	// Output: 9911-2233
}

func ExampleParsePhone() {
	fmt.Println(converter.ParsePhone("+976 9911-2233"))
	// Output: 97699112233
}

func ExampleFormatPhoneInternational() {
	fmt.Println(converter.FormatPhoneInternational("99112233"))
	// Output: +976 9911 2233
}
