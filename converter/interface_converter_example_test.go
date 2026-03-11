package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleInterfaceToBytes() {
	data := map[string]string{"key": "value"}
	bytes, err := converter.InterfaceToBytes(data)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(string(bytes))
	// Output: {"key":"value"}
}

func ExampleInterfaceToInt64() {
	var val interface{} = float64(42) // JSON-оос ирсэн тоо
	result, err := converter.InterfaceToInt64(val, "тоо биш")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(result)
	// Output: 42
}

func ExampleInterfaceToUint() {
	var val interface{} = 100
	result, err := converter.InterfaceToUint(val)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(result)
	// Output: 100
}

func ExampleInterfaceToString() {
	var val interface{} = "Сайн байна уу"
	result, err := converter.InterfaceToString(val, "string биш")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(result)
	// Output: Сайн байна уу
}

func ExampleInterfaceToMap() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "Бат", Age: 25}
	m, err := converter.InterfaceToMap(user)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(m["name"])
	// Output: Бат
}

func ExampleInterfaceToStruct() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := map[string]interface{}{
		"name": "Дорж",
		"age":  30,
	}

	var user User
	err := converter.InterfaceToStruct(data, &user)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(user.Name, user.Age)
	// Output: Дорж 30
}
