package converter

import (
	"bytes"
	"encoding/json"
)

// PrettyJSON нь JSON string-ийг 2 space indent-тэй форматлана.
func PrettyJSON(jsonStr string) (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(jsonStr), "", "  ")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// MinifyJSON нь JSON string-ийг compact болгоно.
func MinifyJSON(jsonStr string) (string, error) {
	var buf bytes.Buffer
	err := json.Compact(&buf, []byte(jsonStr))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// MergeJSON нь олон JSON string-ийг нэгтгэнэ (shallow merge).
func MergeJSON(jsonStrings ...string) (string, error) {
	merged := make(map[string]json.RawMessage)
	for _, s := range jsonStrings {
		var m map[string]json.RawMessage
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			return "", err
		}
		for k, v := range m {
			merged[k] = v
		}
	}
	result, err := json.Marshal(merged)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
