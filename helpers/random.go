package helpers

import (
	"math/rand/v2"
)

// RandInt нь [min, max] хүрээнд санамсаргүй бүхэл тоо буцаана.
func RandInt(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.IntN(max-min+1)
}

// RandFloat нь [min, max) хүрээнд санамсаргүй бутархай тоо буцаана.
func RandFloat(min, max float64) float64 {
	if min >= max {
		return min
	}
	return min + rand.Float64()*(max-min)
}

// RandElement нь slice-аас санамсаргүй элемент буцаана.
func RandElement[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	return slice[rand.IntN(len(slice))]
}

// Shuffle нь slice-ийн элементүүдийг санамсаргүйгээр холино.
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
