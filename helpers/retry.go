package helpers

import "time"

// Retry нь function-ийг амжилтгүй бол дахин оролдоно.
//
// attempts: нийт оролдлогын тоо
// delay: оролдлого хоорондын хүлээх хугацаа
func Retry(attempts int, delay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		if i < attempts-1 {
			time.Sleep(delay)
		}
	}
	return err
}

// RetryWithBackoff нь exponential backoff-тэй дахин оролдоно.
//
// attempts: нийт оролдлогын тоо
// initialDelay: эхний хүлээх хугацаа (дараагийнх 2 дахин)
func RetryWithBackoff(attempts int, initialDelay time.Duration, fn func() error) error {
	var err error
	delay := initialDelay
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		if i < attempts-1 {
			time.Sleep(delay)
			delay *= 2
		}
	}
	return err
}
