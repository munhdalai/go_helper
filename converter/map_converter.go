package converter

import (
	"encoding/json"
	"errors"
)

// MapToMap нь map дотроос key-гээр дотоод map авна.
func MapToMap(m map[string]interface{}, key, errMsg string) (map[string]interface{}, error) {
	result, ok := m[key].(map[string]interface{})
	if !ok {
		return nil, errors.New(errMsg)
	}
	return result, nil
}

// MapToStruct нь map-ийг struct руу JSON-ээр дамжуулан хөрвүүлнэ.
func MapToStruct(m interface{}, v interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// MapToKeyValueSlice нь map-ийг key, value slice болгон хөрвүүлнэ.
// exclude map дахь key-үүдийг алгасна.
func MapToKeyValueSlice(m, exclude map[string]interface{}) (keys, values []string) {
	for key, value := range m {
		if _, ok := exclude[key]; ok {
			continue
		}
		if vs, err := InterfaceToString(value, ""); err == nil {
			keys = append(keys, key)
			values = append(values, vs)
		}
	}
	return keys, values
}
