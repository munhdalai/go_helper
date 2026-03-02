package data

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONRaw нь database-д JSON хадгалах зориулалттай custom type.
type JSONRaw json.RawMessage

func (j JSONRaw) Value() (driver.Value, error) {
	return []byte(j), nil
}

func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	return json.Unmarshal(asBytes, j)
}

func (j JSONRaw) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("JSONRaw: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
