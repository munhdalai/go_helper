package converter

import (
	"fmt"
	"math"
	"strings"
)

// ones нь 1-9 тоонуудын үндсэн хэлбэр.
var ones = []string{
	"", "нэг", "хоёр", "гурав", "дөрөв", "тав", "зургаа", "долоо", "найм", "ес",
}

// onesConnecting нь 1-9 тоонуудын холболтын хэлбэр (дараа нь үг байхад).
var onesConnecting = []string{
	"", "нэг", "хоёр", "гурван", "дөрвөн", "таван", "зургаан", "долоон", "найман", "есөн",
}

// tens нь аравтын үндсэн хэлбэр.
var tens = []string{
	"", "арав", "хорь", "гуч", "дөч", "тавь", "жар", "дал", "ная", "ер",
}

// tensConnecting нь аравтын холболтын хэлбэр.
var tensConnecting = []string{
	"", "арван", "хорин", "гучин", "дөчин", "тавин", "жаран", "далан", "наян", "ерэн",
}

var scales = []struct {
	value int64
	name  string
}{
	{1_000_000_000_000, "их наяд"},
	{1_000_000_000, "тэрбум"},
	{1_000_000, "сая"},
	{1_000, "мянга"},
}

// twoDigits нь 0-99 хүртэлх тоог монгол үгээр буцаана.
func twoDigits(n int64) string {
	if n == 0 {
		return ""
	}
	if n < 10 {
		return ones[n]
	}
	if n == 10 {
		return "арав"
	}
	if n < 20 {
		return "арван " + ones[n-10]
	}
	t := tens[n/10]
	o := ones[n%10]
	if o == "" {
		return t
	}
	return tensConnecting[n/10] + " " + o
}

// threeDigits нь 0-999 хүртэлх тоог монгол үгээр буцаана.
func threeDigits(n int64) string {
	if n == 0 {
		return ""
	}
	h := n / 100
	rest := n % 100
	if h > 0 && rest == 0 {
		return onesConnecting[h] + " зуу"
	}
	if h > 0 {
		return onesConnecting[h] + " зуун " + twoDigits(rest)
	}
	return twoDigits(rest)
}

// MoneyToWords нь мөнгөн дүнг монгол үгээр бичнэ.
//
// Жишээ:
//
//	MoneyToWords(1500)    -> "нэг мянга таван зуун төгрөг"
//	MoneyToWords(1000000) -> "нэг сая төгрөг"
//	MoneyToWords(99.50)   -> "ерэн есөн төгрөг тавин мөнгө"
func MoneyToWords(amount float64) string {
	if amount == 0 {
		return "тэг төгрөг"
	}

	negative := amount < 0
	amount = math.Abs(amount)

	wholePart := int64(amount)
	decimalPart := int64(math.Round((amount - float64(wholePart)) * 100))

	var result string
	if wholePart > 0 {
		result = finalForm(intToWords(wholePart)) + " төгрөг"
	}

	if decimalPart > 0 {
		decWords := finalForm(intToWords(decimalPart))
		if wholePart > 0 {
			result += " " + decWords + " мөнгө"
		} else {
			result = decWords + " мөнгө"
		}
	}

	if negative {
		result = "хасах " + result
	}

	return result
}

// intToWords нь бүхэл тоог монгол үгээр бичнэ.
func intToWords(n int64) string {
	if n == 0 {
		return "тэг"
	}

	var parts []string
	for _, s := range scales {
		if n >= s.value {
			count := n / s.value
			n %= s.value
			countWords := connectingForm(threeDigits(count))
			parts = append(parts, countWords+" "+s.name)
		}
	}

	if n > 0 {
		parts = append(parts, threeDigits(n))
	}

	return strings.Join(parts, " ")
}

// suffixMap нь тоон үгийн холболтын хэлбэрүүд.
var suffixMap = map[string]string{
	"нэг":    "нэг",
	"хоёр":   "хоёр",
	"гурав":  "гурван",
	"дөрөв":  "дөрвөн",
	"тав":    "таван",
	"зургаа": "зургаан",
	"долоо":  "долоон",
	"найм":   "найман",
	"ес":     "есөн",
	"арав":   "арван",
	"хорь":   "хорин",
	"гуч":    "гучин",
	"дөч":    "дөчин",
	"тавь":   "тавин",
	"жар":    "жаран",
	"дал":    "далан",
	"ная":    "наян",
	"ер":     "ерэн",
	"зуу":    "зуун",
	"мянга":  "мянган",
}

// connectingForm нь тоон үгийн сүүлийн хэсгийг холболтын хэлбэрт оруулна.
// "тав" -> "таван", "зуу" -> "зуун", "ер" -> "ерэн" гэх мэт.
// "нэг" нь хэвээрээ үлдэнэ (мянга, сая-ын өмнө "нэг" байдаг).
func connectingForm(s string) string {
	return replaceLastWord(s, suffixMap)
}

// finalForm нь төгрөг/мөнгө-ийн өмнөх тоон үгийг зөв хэлбэрт оруулна.
// "нэг" -> "нэгэн", "тав" -> "таван" гэх мэт.
func finalForm(s string) string {
	finalMap := make(map[string]string, len(suffixMap))
	for k, v := range suffixMap {
		finalMap[k] = v
	}
	finalMap["нэг"] = "нэгэн"
	return replaceLastWord(s, finalMap)
}

// replaceLastWord нь string-ийн сүүлийн үгийг map-аас солино.
func replaceLastWord(s string, m map[string]string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	last := words[len(words)-1]
	if replacement, ok := m[last]; ok {
		words[len(words)-1] = replacement
	}

	return strings.Join(words, " ")
}

// FormatMoney нь тоог мөнгөний формат руу хөрвүүлнэ (мянгатын таслалтай).
//
// Жишээ:
//
//	FormatMoney(1234567.89) -> "1,234,567.89"
//	FormatMoney(1000)       -> "1,000.00"
//	FormatMoney(99.5)       -> "99.50"
func FormatMoney(amount float64) string {
	negative := amount < 0
	amount = math.Abs(amount)

	wholePart := int64(amount)
	decimalPart := int64(math.Round((amount - float64(wholePart)) * 100))

	wholeStr := Int64ToString(wholePart)

	// Мянгатын таслал нэмэх
	n := len(wholeStr)
	if n > 3 {
		var groups []string
		for n > 0 {
			start := n - 3
			if start < 0 {
				start = 0
			}
			groups = append([]string{wholeStr[start:n]}, groups...)
			n = start
		}
		wholeStr = strings.Join(groups, ",")
	}

	result := fmt.Sprintf("%s.%02d", wholeStr, decimalPart)
	if negative {
		result = "-" + result
	}
	return result
}

// FormatTugrik нь тоог төгрөгийн формат руу хөрвүүлнэ.
//
// Жишээ:
//
//	FormatTugrik(1234567.89) -> "₮1,234,567.89"
//	FormatTugrik(1000)       -> "₮1,000.00"
func FormatTugrik(amount float64) string {
	if amount < 0 {
		return "-₮" + FormatMoney(math.Abs(amount))
	}
	return "₮" + FormatMoney(amount)
}

// ParseMoney нь мөнгөний форматтай string-ийг float64 руу хөрвүүлнэ.
//
// Жишээ:
//
//	ParseMoney("1,234,567.89") -> 1234567.89
//	ParseMoney("₮1,000.00")   -> 1000.00
//	ParseMoney("1000")         -> 1000.00
func ParseMoney(s string) float64 {
	s = strings.TrimSpace(s)

	negative := false
	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	}

	s = strings.TrimPrefix(s, "₮")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSpace(s)

	result := StringToFloat64(s)
	if negative {
		return -result
	}
	return result
}

// MongoCentsToTugrik нь мөнгөг (100 мөнгө = 1 төгрөг) төгрөг рүү хөрвүүлнэ.
//
// Жишээ:
//
//	MongoCentsToTugrik(15050) -> 150.50
//	MongoCentsToTugrik(100)   -> 1.00
func MongoCentsToTugrik(mongo int64) float64 {
	return float64(mongo) / 100
}

// TugrikToMongoCents нь төгрөгийг мөнгө рүү хөрвүүлнэ (1 төгрөг = 100 мөнгө).
//
// Жишээ:
//
//	TugrikToMongoCents(150.50) -> 15050
//	TugrikToMongoCents(1.00)   -> 100
func TugrikToMongoCents(tugrik float64) int64 {
	return int64(math.Round(tugrik * 100))
}
