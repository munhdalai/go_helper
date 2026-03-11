package helpers

import (
	"fmt"
	"math"
	"time"
)

// StartOfDay нь өгөгдсөн өдрийн эхлэлийг буцаана (00:00:00).
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay нь өгөгдсөн өдрийн төгсгөлийг буцаана (23:59:59.999999999).
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// DaysBetween нь хоёр цагийн хоорондох өдрийн тоог буцаана.
func DaysBetween(a, b time.Time) int {
	a = StartOfDay(a)
	b = StartOfDay(b)
	diff := b.Sub(a).Hours() / 24
	return int(math.Abs(diff))
}

// IsWeekend нь амралтын өдөр эсэхийг шалгана.
func IsWeekend(t time.Time) bool {
	day := t.Weekday()
	return day == time.Saturday || day == time.Sunday
}

// IsToday нь өнөөдөр эсэхийг шалгана.
func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// IsFuture нь ирээдүйн цаг эсэхийг шалгана.
func IsFuture(t time.Time) bool {
	return t.After(time.Now())
}

// IsPast нь өнгөрсөн цаг эсэхийг шалгана.
func IsPast(t time.Time) bool {
	return t.Before(time.Now())
}

// RelativeTime нь цагийг харьцангуй текстээр буцаана (Монгол хэлээр).
func RelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)
	future := diff < 0
	if future {
		diff = -diff
	}

	suffix := "өмнө"
	if future {
		suffix = "дараа"
	}

	seconds := int(diff.Seconds())
	minutes := int(diff.Minutes())
	hours := int(diff.Hours())
	days := int(diff.Hours() / 24)

	switch {
	case seconds < 60:
		return fmt.Sprintf("%d секундын %s", seconds, suffix)
	case minutes < 60:
		return fmt.Sprintf("%d минутын %s", minutes, suffix)
	case hours < 24:
		return fmt.Sprintf("%d цагийн %s", hours, suffix)
	case days < 30:
		return fmt.Sprintf("%d хоногийн %s", days, suffix)
	case days < 365:
		return fmt.Sprintf("%d сарын %s", days/30, suffix)
	default:
		return fmt.Sprintf("%d жилийн %s", days/365, suffix)
	}
}

// AddBusinessDays нь ажлын өдрүүд нэмнэ (Бямба, Ням-ыг алгасана).
func AddBusinessDays(t time.Time, days int) time.Time {
	direction := 1
	if days < 0 {
		direction = -1
		days = -days
	}
	added := 0
	for added < days {
		t = t.AddDate(0, 0, direction)
		if !IsWeekend(t) {
			added++
		}
	}
	return t
}
