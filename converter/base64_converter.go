package converter

import "encoding/base64"

// Base64Encode нь string-ийг base64 руу encode хийнэ.
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode нь base64 string-ийг decode хийнэ.
func Base64Decode(encoded string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Base64EncodeBytes нь []byte-ийг base64 string руу encode хийнэ.
func Base64EncodeBytes(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64DecodeBytes нь base64 string-ийг []byte руу decode хийнэ.
func Base64DecodeBytes(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
