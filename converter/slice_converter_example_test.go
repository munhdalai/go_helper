package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleStringSliceToIntSlice() {
	strs := []string{"1", "2", "3"}
	ints := converter.StringSliceToIntSlice(strs)
	fmt.Println(ints)
	// Output: [1 2 3]
}

func ExampleIntSliceToStringSlice() {
	ints := []int{10, 20, 30}
	strs := converter.IntSliceToStringSlice(ints)
	fmt.Println(strs)
	// Output: [10 20 30]
}

func ExampleSliceToMap() {
	type User struct {
		ID   int
		Name string
	}

	users := []User{
		{ID: 1, Name: "Бат"},
		{ID: 2, Name: "Дорж"},
	}

	m := converter.SliceToMap(users, func(u User) int {
		return u.ID
	})

	fmt.Println(m[1].Name)
	fmt.Println(m[2].Name)
	// Output:
	// Бат
	// Дорж
}
