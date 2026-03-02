package converter

import "encoding/json"

// StructToMap нь struct-ийг map[string]interface{} руу хөрвүүлнэ.
func StructToMap(v interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}

// StructToJSON нь struct-ийг JSON string руу хөрвүүлнэ.
func StructToJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// JSONToStruct нь JSON string-ийг struct руу хөрвүүлнэ.
func JSONToStruct(jsonStr string, target interface{}) error {
	return json.Unmarshal([]byte(jsonStr), target)
}
