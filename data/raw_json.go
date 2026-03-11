package data

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONRaw нь database-д JSON хадгалах зориулалттай custom type.
type JSONRaw json.RawMessage

// Value нь database/sql драйверт []byte утга буцаана.
func (j JSONRaw) Value() (driver.Value, error) {
	return []byte(j), nil
}

// Scan нь database-ийн []byte утгыг JSONRaw руу уншина.
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	return json.Unmarshal(asBytes, j)
}

// MarshalJSON нь JSONRaw-ийг JSON болгон буцаана. nil бол "null" буцаана.
func (j JSONRaw) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON нь JSON өгөгдлийг JSONRaw руу хадгална.
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("JSONRaw: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
