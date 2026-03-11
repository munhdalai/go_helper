package crypto_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/crypto"
)

func ExampleAESEncrypt() {
	key := []byte("0123456789abcdef0123456789abcdef") // 32 bytes = AES-256
	plaintext := []byte("Hello, World!")
	encrypted, err := crypto.AESEncrypt(plaintext, key)
	fmt.Println(err)
	fmt.Println(len(encrypted) > 0)
	// Output:
	// <nil>
	// true
}

func ExampleAESDecrypt() {
	key := []byte("0123456789abcdef0123456789abcdef")
	plaintext := []byte("Hello, World!")
	encrypted, _ := crypto.AESEncrypt(plaintext, key)
	decrypted, err := crypto.AESDecrypt(encrypted, key)
	fmt.Println(string(decrypted), err)
	// Output: Hello, World! <nil>
}

func ExampleAESEncryptString() {
	encrypted, err := crypto.AESEncryptString("Hello, World!", "my-secret-key")
	fmt.Println(err)
	fmt.Println(len(encrypted) > 0)
	// Output:
	// <nil>
	// true
}

func ExampleAESDecryptString() {
	encrypted, _ := crypto.AESEncryptString("Hello, World!", "my-secret-key")
	decrypted, err := crypto.AESDecryptString(encrypted, "my-secret-key")
	fmt.Println(decrypted, err)
	// Output: Hello, World! <nil>
}

func ExampleHMACSign() {
	signature := crypto.HMACSign([]byte("hello"), []byte("secret"))
	fmt.Println(len(signature))
	// Output: 64
}

func ExampleHMACVerify() {
	key := []byte("secret")
	message := []byte("hello")
	signature := crypto.HMACSign(message, key)
	fmt.Println(crypto.HMACVerify(message, key, signature))
	fmt.Println(crypto.HMACVerify(message, key, "invalid"))
	// Output:
	// true
	// false
}
