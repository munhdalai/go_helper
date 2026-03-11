package converter_test

import (
	"fmt"
	"sort"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleMapToMap() {
	m := map[string]interface{}{
		"user": map[string]interface{}{
			"name": "Бат",
			"age":  25,
		},
	}

	user, err := converter.MapToMap(m, "user", "user олдсонгүй")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(user["name"])
	// Output: Бат
}

func ExampleMapToStruct() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	m := map[string]interface{}{
		"name": "Дорж",
		"age":  30,
	}

	var user User
	err := converter.MapToStruct(m, &user)
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(user.Name, user.Age)
	// Output: Дорж 30
}

func ExampleMapToKeyValueSlice() {
	m := map[string]interface{}{
		"name": "Бат",
		"city": "УБ",
		"id":   "123",
	}
	exclude := map[string]interface{}{
		"id": nil,
	}

	keys, values := converter.MapToKeyValueSlice(m, exclude)

	// map-ийн дараалал тодорхойгүй тул sort хийж харуулна
	paired := make([]string, len(keys))
	for i := range keys {
		paired[i] = keys[i] + "=" + values[i]
	}
	sort.Strings(paired)
	fmt.Println(paired)
	// Output: [city=УБ name=Бат]
}
