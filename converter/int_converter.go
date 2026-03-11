package converter

import "strconv"

// UintToString нь uint-ийг string руу хөрвүүлнэ.
func UintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

// IntToString нь int-ийг string руу хөрвүүлнэ.
func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

// Int64ToString нь int64-ийг string руу хөрвүүлнэ.
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
