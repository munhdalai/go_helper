package validate

import (
	"net"
	"net/url"
	"regexp"
	"unicode"
)

var (
	bankAccountRegex   = regexp.MustCompile(`^\d{10,12}$`)
	mongolianNameRegex = regexp.MustCompile(`^[а-яА-ЯөӨүҮёЁ\s\-\.]{2,50}$`)
)

// IsURL нь зөв URL эсэхийг шалгана.
func IsURL(val string) bool {
	u, err := url.ParseRequestURI(val)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

// IsIP нь зөв IP хаяг эсэхийг шалгана.
func IsIP(val string) bool {
	return net.ParseIP(val) != nil
}

// IsStrongPassword нь нууц үг хүчтэй эсэхийг шалгана.
// Дор хаяж 8 тэмдэгт, том/жижиг үсэг, тоо, тусгай тэмдэгт агуулсан байх ёстой.
func IsStrongPassword(val string) bool {
	if len(val) < 8 {
		return false
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, c := range val {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}

// IsBankAccount нь банкны дансны дугаар зөв эсэхийг шалгана (10-12 оронтой тоо).
func IsBankAccount(val string) bool {
	return bankAccountRegex.MatchString(val)
}

// IsMongolianName нь Монгол нэр зөв эсэхийг шалгана (Кирилл, 2-50 тэмдэгт).
func IsMongolianName(val string) bool {
	return mongolianNameRegex.MatchString(val)
}
