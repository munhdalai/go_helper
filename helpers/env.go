package helpers

import "os"

// Env нь environment variable уншина. Утга байхгүй бол defaultVal буцаана.
func Env(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// Ternary нь нөхцөлт утга сонгох helper.
// Ternary(true, "a", "b") -> "a"
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
