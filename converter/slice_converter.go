package converter

import "strconv"

// StringSliceToIntSlice нь []string-ийг []int руу хөрвүүлнэ.
// Хөрвүүлж чадаагүй утгыг 0 болгоно.
func StringSliceToIntSlice(slice []string) []int {
	result := make([]int, len(slice))
	for i, s := range slice {
		result[i] = StringToInt(s)
	}
	return result
}

// IntSliceToStringSlice нь []int-ийг []string руу хөрвүүлнэ.
func IntSliceToStringSlice(slice []int) []string {
	result := make([]string, len(slice))
	for i, n := range slice {
		result[i] = strconv.Itoa(n)
	}
	return result
}

// SliceToMap нь slice-ийг map болгоно. keyFunc нь элемент бүрээс key гаргана.
func SliceToMap[T any, K comparable](slice []T, keyFunc func(T) K) map[K]T {
	result := make(map[K]T, len(slice))
	for _, item := range slice {
		result[keyFunc(item)] = item
	}
	return result
}
