package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify нь string-ийг URL-д тохиромжтой slug болгоно.
// "Hello World!" -> "hello-world"
func Slugify(s string) string {
	slug := strings.ToLower(strings.TrimSpace(s))
	slug = nonAlphanumeric.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}

// Truncate нь string-ийг заасан урттай болгож тасалаад "..." нэмнэ.
func Truncate(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

// Capitalize нь эхний үсгийг том болгоно.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// ToCamelCase нь snake_case string-ийг CamelCase руу хөрвүүлнэ.
func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		result.WriteString(Capitalize(strings.ToLower(part)))
	}
	return result.String()
}
