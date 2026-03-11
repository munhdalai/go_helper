package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleMoneyToWords() {
	fmt.Println(converter.MoneyToWords(0))
	fmt.Println(converter.MoneyToWords(1))
	fmt.Println(converter.MoneyToWords(10))
	fmt.Println(converter.MoneyToWords(100))
	fmt.Println(converter.MoneyToWords(1500))
	fmt.Println(converter.MoneyToWords(52000))
	fmt.Println(converter.MoneyToWords(1000000))
	fmt.Println(converter.MoneyToWords(99.50))
	// Output:
	// тэг төгрөг
	// нэгэн төгрөг
	// арван төгрөг
	// нэг зуун төгрөг
	// нэг мянга таван зуун төгрөг
	// тавин хоёр мянган төгрөг
	// нэг сая төгрөг
	// ерэн есөн төгрөг тавин мөнгө
}

func ExampleMoneyToWords_large() {
	fmt.Println(converter.MoneyToWords(1234567))
	fmt.Println(converter.MoneyToWords(5000000000))
	// Output:
	// нэг сая хоёр зуун гучин дөрвөн мянга таван зуун жаран долоон төгрөг
	// таван тэрбум төгрөг
}

func ExampleMoneyToWords_negative() {
	fmt.Println(converter.MoneyToWords(-5000))
	// Output: хасах таван мянган төгрөг
}

func ExampleFormatMoney() {
	fmt.Println(converter.FormatMoney(1234567.89))
	fmt.Println(converter.FormatMoney(1000))
	fmt.Println(converter.FormatMoney(99.5))
	fmt.Println(converter.FormatMoney(0))
	// Output:
	// 1,234,567.89
	// 1,000.00
	// 99.50
	// 0.00
}

func ExampleFormatMoney_negative() {
	fmt.Println(converter.FormatMoney(-1500.75))
	// Output: -1,500.75
}

func ExampleFormatTugrik() {
	fmt.Println(converter.FormatTugrik(1234567.89))
	fmt.Println(converter.FormatTugrik(1000))
	fmt.Println(converter.FormatTugrik(-500))
	// Output:
	// ₮1,234,567.89
	// ₮1,000.00
	// -₮500.00
}

func ExampleParseMoney() {
	fmt.Println(converter.ParseMoney("1,234,567.89"))
	fmt.Println(converter.ParseMoney("₮1,000.00"))
	fmt.Println(converter.ParseMoney("1000"))
	fmt.Println(converter.ParseMoney("99.50"))
	// Output:
	// 1.23456789e+06
	// 1000
	// 1000
	// 99.5
}

func ExampleMongoCentsToTugrik() {
	fmt.Println(converter.MongoCentsToTugrik(15050))
	fmt.Println(converter.MongoCentsToTugrik(100))
	fmt.Println(converter.MongoCentsToTugrik(0))
	// Output:
	// 150.5
	// 1
	// 0
}

func ExampleTugrikToMongoCents() {
	fmt.Println(converter.TugrikToMongoCents(150.50))
	fmt.Println(converter.TugrikToMongoCents(1.00))
	fmt.Println(converter.TugrikToMongoCents(0))
	// Output:
	// 15050
	// 100
	// 0
}
