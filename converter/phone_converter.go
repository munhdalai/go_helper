package converter

import (
	"regexp"
	"strings"
)

var nonDigitRegex = regexp.MustCompile(`\D`)

// FormatPhone нь утасны дугаарыг "9911-2233" форматаар буцаана.
func FormatPhone(phone string) string {
	digits := nonDigitRegex.ReplaceAllString(phone, "")
	if len(digits) == 8 {
		return digits[:4] + "-" + digits[4:]
	}
	return phone
}

// ParsePhone нь утасны дугаараас тоон бус тэмдэгтүүдийг арилгана.
func ParsePhone(phone string) string {
	return nonDigitRegex.ReplaceAllString(phone, "")
}

// FormatPhoneInternational нь утасны дугаарыг "+976 9911 2233" форматаар буцаана.
func FormatPhoneInternational(phone string) string {
	digits := nonDigitRegex.ReplaceAllString(phone, "")
	// Remove leading 976 if present
	digits = strings.TrimPrefix(digits, "976")
	if len(digits) == 8 {
		return "+976 " + digits[:4] + " " + digits[4:]
	}
	return phone
}
