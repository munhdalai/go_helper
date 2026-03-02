package converter

import (
	"encoding/json"
	"errors"
	"fmt"
)

// InterfaceToBytes нь interface{}-ийг JSON bytes руу хөрвүүлнэ.
func InterfaceToBytes(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}

// InterfaceToInt64 нь interface{}-ийг int64 руу хөрвүүлнэ.
func InterfaceToInt64(val interface{}, errMsg string) (int64, error) {
	switch v := val.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case json.Number:
		return v.Int64()
	default:
		return 0, errors.New(errMsg)
	}
}

// InterfaceToUint нь interface{}-ийг uint руу хөрвүүлнэ.
func InterfaceToUint(val interface{}) (uint, error) {
	switch v := val.(type) {
	case uint:
		return v, nil
	case int:
		if v < 0 {
			return 0, fmt.Errorf("cannot convert negative int to uint")
		}
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("cannot convert negative float64 to uint")
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, fmt.Errorf("cannot convert negative int64 to uint")
		}
		return uint(v), nil
	case string:
		return StringToUint(v), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", val)
	}
}

// InterfaceToString нь interface{}-ийг string руу хөрвүүлнэ.
func InterfaceToString(val interface{}, errMsg string) (string, error) {
	result, ok := val.(string)
	if !ok {
		return "", errors.New(errMsg)
	}
	return result, nil
}

// InterfaceToMap нь interface{}-ийг map[string]interface{} руу хөрвүүлнэ.
func InterfaceToMap(val interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}

// InterfaceToStruct нь interface{}-ийг заасан struct руу хөрвүүлнэ.
func InterfaceToStruct(val interface{}, target interface{}) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}
