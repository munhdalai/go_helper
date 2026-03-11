package data

import (
	"database/sql/driver"
	"encoding/json"
)

// Null нь nullable утгыг илэрхийлэх generic төрөл.
type Null[T any] struct {
	Val   T
	Valid bool
}

// NewNull нь шинэ Null утга үүсгэнэ (Valid=true).
func NewNull[T any](val T) Null[T] {
	return Null[T]{Val: val, Valid: true}
}

// NullFrom нь pointer-ээс Null утга үүсгэнэ.
func NullFrom[T any](ptr *T) Null[T] {
	if ptr == nil {
		return Null[T]{}
	}
	return Null[T]{Val: *ptr, Valid: true}
}

// IsNull нь утга null эсэхийг шалгана.
func (n Null[T]) IsNull() bool {
	return !n.Valid
}

// ValueOrDefault нь утга байвал буцаана, байхгүй бол default утга буцаана.
func (n Null[T]) ValueOrDefault(defaultVal T) T {
	if !n.Valid {
		return defaultVal
	}
	return n.Val
}

// Ptr нь утгын pointer буцаана, null бол nil буцаана.
func (n Null[T]) Ptr() *T {
	if !n.Valid {
		return nil
	}
	v := n.Val
	return &v
}

// MarshalJSON нь JSON руу хөрвүүлнэ.
func (n Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

// UnmarshalJSON нь JSON-оос уншина.
func (n *Null[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.Val)
}

// Value нь database/sql/driver.Value interface хэрэгжүүлнэ.
func (n Null[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	data, err := json.Marshal(n.Val)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

// Scan нь database/sql Scanner interface хэрэгжүүлнэ.
func (n *Null[T]) Scan(src interface{}) error {
	if src == nil {
		n.Valid = false
		return nil
	}
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		var err error
		data, err = json.Marshal(src)
		if err != nil {
			return err
		}
	}
	n.Valid = true
	return json.Unmarshal(data, &n.Val)
}
