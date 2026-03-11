package helpers_test

import (
	"fmt"
	"os"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleEnv() {
	os.Setenv("APP_PORT", "8080")

	fmt.Println(helpers.Env("APP_PORT", "3000"))        // тохируулсан утга
	fmt.Println(helpers.Env("UNKNOWN_KEY", "default"))   // default утга
	// Output:
	// 8080
	// default
}

func ExampleTernary() {
	age := 20

	status := helpers.Ternary(age >= 18, "насанд хүрсэн", "хүүхэд")
	fmt.Println(status)

	max := helpers.Ternary(10 > 5, 10, 5)
	fmt.Println(max)
	// Output:
	// насанд хүрсэн
	// 10
}
