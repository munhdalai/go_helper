package helpers_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleRetry() {
	attempt := 0
	err := helpers.Retry(3, 100*time.Millisecond, func() error {
		attempt++
		if attempt < 3 {
			return errors.New("амжилтгүй")
		}
		return nil // 3 дахь оролдлогод амжилттай
	})

	fmt.Println(err)
	fmt.Println(attempt)
	// Output:
	// <nil>
	// 3
}

func ExampleRetryWithBackoff() {
	attempt := 0
	err := helpers.RetryWithBackoff(3, 50*time.Millisecond, func() error {
		attempt++
		if attempt < 2 {
			return errors.New("амжилтгүй")
		}
		return nil // 2 дахь оролдлогод амжилттай
	})

	fmt.Println(err)
	fmt.Println(attempt)
	// Output:
	// <nil>
	// 2
}
