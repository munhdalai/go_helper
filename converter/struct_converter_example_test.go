package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleStructToMap() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "Бат", Age: 25}
	m, err := converter.StructToMap(user)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	// Output:
	// Бат
	// 25
}

func ExampleStructToJSON() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{Name: "Бат", Age: 25}
	jsonStr, err := converter.StructToJSON(user)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(jsonStr)
	// Output: {"name":"Бат","age":25}
}

func ExampleJSONToStruct() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	jsonStr := `{"name":"Дорж","age":30}`
	var user User
	err := converter.JSONToStruct(jsonStr, &user)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(user.Name, user.Age)
	// Output: Дорж 30
}
