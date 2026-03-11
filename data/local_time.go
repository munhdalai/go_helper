package data

import (
	"database/sql/driver"
	"time"
)

// LocalTime нь огноо + цагийг хадгалах custom type.
type LocalTime time.Time

const localTimeFormat = "2006-01-02 15:04:05"

// UnmarshalJSON нь JSON string-ийг LocalTime руу parse хийнэ.
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(`"`+localTimeFormat+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = LocalTime(parsed)
	return nil
}

// MarshalJSON нь LocalTime-ийг JSON string руу хөрвүүлнэ.
func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localTimeFormat)+2)
	b = append(b, '"')
	b = append(b, []byte(t.String())...)
	b = append(b, '"')
	return b, nil
}

// String нь LocalTime-ийг "2006-01-02 15:04:05" формат руу хөрвүүлнэ.
func (t LocalTime) String() string {
	if time.Time(t).IsZero() {
		return "0000-00-00 00:00:00"
	}
	return time.Time(t).Format(localTimeFormat)
}

// Value нь database/sql драйверт утга буцаана.
func (t LocalTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return time.Now(), nil
	}
	return time.Time(t), nil
}

// Scan нь database-ийн утгыг LocalTime руу уншина.
func (t *LocalTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = LocalTime(vt)
	case string:
		parsed, _ := time.Parse("2006-01-02 15:04:05", vt)
		*t = LocalTime(parsed)
	}
	return nil
}

// ToTime нь LocalTime-ийг time.Time руу хөрвүүлнэ.
func (t LocalTime) ToTime() time.Time {
	return time.Time(t)
}
