package helpers

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"math"
	"math/big"
	"regexp"
	"strings"
)

// GeneratePassword нь SHA1 hash үүсгэнэ.
func GeneratePassword(password string) string {
	hash := sha1.Sum([]byte(strings.ToUpper(password)))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase нь CamelCase string-ийг snake_case руу хөрвүүлнэ.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandString нь n урттай санамсаргүй тэмдэгт мөр үүсгэнэ.
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[idx.Int64()]
	}
	return string(b)
}

// StringInSlice нь string slice дотор утга байгаа эсэхийг шалгана.
func StringInSlice(target string, list []string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

// GenerateOTP нь заасан урттай нэг удаагын нууц үг үүсгэнэ.
func GenerateOTP(length int) uint {
	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length)) - 1
	n, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return uint(n.Int64() + min)
}

// GenerateMD5Hash нь MD5 hash үүсгэнэ.
func GenerateMD5Hash(values ...string) string {
	combined := strings.ToLower(strings.TrimSpace(strings.Join(values, "")))
	hash := md5.Sum([]byte(combined))
	return hex.EncodeToString(hash[:])
}

// UniqueIntSlice нь int slice-аас давхардлыг арилгана.
func UniqueIntSlice(slice []int) []int {
	seen := make(map[int]struct{})
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// UniqueStringSlice нь string slice-аас давхардлыг арилгана.
func UniqueStringSlice(slice []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// RoundFloat нь float64-ийг 2 орон хүртэл тоймлоно.
func RoundFloat(v float64) float64 {
	return math.Floor(v*100) / 100
}

// Mask нь string-ийн тодорхой тэмдэгтүүдийг * болгоно.
func Mask(str string, maskType int) string {
	if maskType == 1 {
		var result strings.Builder
		for i, c := range str {
			if i%2 == 1 {
				result.WriteRune('*')
			} else {
				result.WriteRune(c)
			}
		}
		return result.String()
	}
	return str
}
