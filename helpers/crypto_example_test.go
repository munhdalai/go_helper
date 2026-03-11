package helpers_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleGenerateUUID() {
	uuid := helpers.GenerateUUID()
	fmt.Println(len(uuid)) // UUID нь 36 тэмдэгт (8-4-4-4-12)
	// Output: 36
}

func ExampleHashPassword() {
	hash, err := helpers.HashPassword("mypassword123")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	// bcrypt hash нь $2a$ эсвэл $2b$-ээр эхэлнэ
	fmt.Println(hash[:4])
	// Output: $2a$
}

func ExampleCheckPassword() {
	hash, _ := helpers.HashPassword("secret123")

	fmt.Println(helpers.CheckPassword("secret123", hash))  // зөв password
	fmt.Println(helpers.CheckPassword("wrongpass", hash))   // буруу password
	// Output:
	// true
	// false
}
