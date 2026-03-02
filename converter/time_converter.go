package converter

import "time"

// TimeToDateString нь time.Time-ийг "2006-01-02" формат руу хөрвүүлнэ.
func TimeToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

// TimeToDateTimeString нь time.Time-ийг "2006-01-02 15:04:05" формат руу хөрвүүлнэ.
func TimeToDateTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// TimeToTimeString нь time.Time-ийг "15:04:05" формат руу хөрвүүлнэ.
func TimeToTimeString(t time.Time) string {
	return t.Format("15:04:05")
}
