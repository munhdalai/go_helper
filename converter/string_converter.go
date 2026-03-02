package converter

import (
	"encoding/json"
	"strconv"
	"time"
)

// StringToInt нь string-ийг int руу хөрвүүлнэ. Алдаа гарвал 0 буцаана.
func StringToInt(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		return 0
	}
	return res
}

// StringToUint нь string-ийг uint руу хөрвүүлнэ.
func StringToUint(num string) uint {
	return uint(StringToInt(num))
}

// StringToInt64 нь string-ийг int64 руу хөрвүүлнэ.
func StringToInt64(num string) int64 {
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

// StringToFloat64 нь string-ийг float64 руу хөрвүүлнэ.
func StringToFloat64(num string) float64 {
	res, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0
	}
	return res
}

// StringToMap нь JSON string-ийг map[string]interface{} руу хөрвүүлнэ.
func StringToMap(jsonString string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	return result, err
}

// StringToMapArr нь JSON string-ийг []map[string]interface{} руу хөрвүүлнэ.
func StringToMapArr(jsonString string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	return result, err
}

// DateStringToTime нь "2006-01-02" форматын string-ийг time.Time руу хөрвүүлнэ.
func DateStringToTime(t string) time.Time {
	c, err := time.Parse("2006-01-02", t)
	if err != nil {
		return time.Time{}
	}
	return c
}

// DateTimeStringToTime нь "2006-01-02 15:04:05" форматын string-ийг time.Time руу хөрвүүлнэ.
func DateTimeStringToTime(t string) time.Time {
	c, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Time{}
	}
	return c
}

// StringToBool нь string-ийг bool руу хөрвүүлнэ.
func StringToBool(s string) bool {
	res, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return res
}
