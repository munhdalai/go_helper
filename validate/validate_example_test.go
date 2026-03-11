package validate_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/validate"
)

func ExampleStruct() {
	type User struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
		Age   int    `validate:"gte=0,lte=150"`
	}

	// Зөв өгөгдөл
	user := User{Name: "Бат", Email: "bat@example.com", Age: 25}
	err := validate.Struct(user)
	fmt.Println(err)

	// Буруу өгөгдөл
	badUser := User{Name: "", Email: "invalid", Age: -1}
	err = validate.Struct(badUser)
	fmt.Println(err != nil)
	// Output:
	// <nil>
	// true
}

func ExampleIsEmail() {
	fmt.Println(validate.IsEmail("user@example.com"))
	fmt.Println(validate.IsEmail("invalid-email"))
	fmt.Println(validate.IsEmail("test@domain.mn"))
	// Output:
	// true
	// false
	// true
}

func ExampleIsPhoneNo() {
	fmt.Println(validate.IsPhoneNo("99112233"))  // 8 оронтой
	fmt.Println(validate.IsPhoneNo("1234567"))   // 7 оронтой - буруу
	fmt.Println(validate.IsPhoneNo("123456789")) // 9 оронтой - буруу
	// Output:
	// true
	// false
	// false
}

func ExampleIsRegNo() {
	fmt.Println(validate.IsRegNo("АБ12345678"))  // 2 кирилл + 8 тоо
	fmt.Println(validate.IsRegNo("AB12345678"))   // латин үсэг - буруу
	fmt.Println(validate.IsRegNo("А1234567"))     // 1 кирилл - буруу
	// Output:
	// true
	// false
	// false
}

func ExampleIsPlateNo() {
	fmt.Println(validate.IsPlateNo("1234АБ"))    // 4 тоо + 2 кирилл
	fmt.Println(validate.IsPlateNo("1234АБВ"))   // 4 тоо + 3 кирилл
	fmt.Println(validate.IsPlateNo("123АБ"))     // 3 тоо - буруу
	// Output:
	// true
	// true
	// false
}

func ExampleIsNumeric() {
	fmt.Println(validate.IsNumeric("12345"))
	fmt.Println(validate.IsNumeric("12.34"))  // бутархай - буруу
	fmt.Println(validate.IsNumeric("abc"))
	// Output:
	// true
	// false
	// false
}
