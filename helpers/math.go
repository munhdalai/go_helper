package helpers

import "cmp"

// Number нь бүх тоон төрлүүдийг илэрхийлэх interface.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Min нь хоёр утгаас бага утгыг буцаана.
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max нь хоёр утгаас их утгыг буцаана.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp нь утгыг [min, max] хязгаарт оруулна.
func Clamp[T cmp.Ordered](val, lo, hi T) T {
	if val < lo {
		return lo
	}
	if val > hi {
		return hi
	}
	return val
}

// Sum нь slice-ийн бүх элементүүдийн нийлбэрийг буцаана.
func Sum[T Number](slice []T) T {
	var total T
	for _, v := range slice {
		total += v
	}
	return total
}

// Average нь slice-ийн дундаж утгыг буцаана.
func Average[T Number](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}
	var total float64
	for _, v := range slice {
		total += float64(v)
	}
	return total / float64(len(slice))
}

// Abs нь тооны абсолют утгыг буцаана.
func Abs[T Number](val T) T {
	if val < 0 {
		return -val
	}
	return val
}
