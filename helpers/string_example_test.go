package helpers_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleSlugify() {
	fmt.Println(helpers.Slugify("Hello World!"))
	fmt.Println(helpers.Slugify("  Go  Programming  "))
	fmt.Println(helpers.Slugify("My Blog Post #1"))
	// Output:
	// hello-world
	// go-programming
	// my-blog-post-1
}

func ExampleTruncate() {
	fmt.Println(helpers.Truncate("Hello, World!", 10))
	fmt.Println(helpers.Truncate("Short", 10))
	// Output:
	// Hello, ...
	// Short
}

func ExampleCapitalize() {
	fmt.Println(helpers.Capitalize("hello"))
	fmt.Println(helpers.Capitalize("world"))
	// Output:
	// Hello
	// World
}

func ExampleToCamelCase() {
	fmt.Println(helpers.ToCamelCase("hello_world"))
	fmt.Println(helpers.ToCamelCase("user_first_name"))
	fmt.Println(helpers.ToCamelCase("created_at"))
	// Output:
	// HelloWorld
	// UserFirstName
	// CreatedAt
}
