package converter_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/converter"
)

func ExampleBase64Encode() {
	encoded := converter.Base64Encode("Hello, World!")
	fmt.Println(encoded)
	// Output: SGVsbG8sIFdvcmxkIQ==
}

func ExampleBase64Decode() {
	decoded, err := converter.Base64Decode("SGVsbG8sIFdvcmxkIQ==")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(decoded)
	// Output: Hello, World!
}

func ExampleBase64EncodeBytes() {
	data := []byte{72, 101, 108, 108, 111}
	encoded := converter.Base64EncodeBytes(data)
	fmt.Println(encoded)
	// Output: SGVsbG8=
}

func ExampleBase64DecodeBytes() {
	data, err := converter.Base64DecodeBytes("SGVsbG8=")
	if err != nil {
		fmt.Println("Алдаа:", err)
		return
	}
	fmt.Println(string(data))
	// Output: Hello
}
