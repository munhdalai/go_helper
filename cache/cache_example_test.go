package cache_test

import (
	"fmt"
	"time"

	"github.com/munhdalai/go_helper/cache"
)

func ExampleNew() {
	c := cache.New[string, int](5 * time.Minute)
	c.Set("age", 25)
	val, ok := c.Get("age")
	fmt.Println(val, ok)
	// Output: 25 true
}

func ExampleCache_Set() {
	c := cache.New[string, string](0)
	c.Set("name", "test")
	val, ok := c.Get("name")
	fmt.Println(val, ok)
	// Output: test true
}

func ExampleCache_SetWithTTL() {
	c := cache.New[string, string](0)
	c.SetWithTTL("key", "value", 1*time.Hour)
	val, ok := c.Get("key")
	fmt.Println(val, ok)
	// Output: value true
}

func ExampleCache_Get() {
	c := cache.New[string, int](0)
	c.Set("a", 1)
	val, ok := c.Get("a")
	fmt.Println(val, ok)
	_, ok = c.Get("b")
	fmt.Println(ok)
	// Output:
	// 1 true
	// false
}

func ExampleCache_Delete() {
	c := cache.New[string, int](0)
	c.Set("a", 1)
	c.Delete("a")
	_, ok := c.Get("a")
	fmt.Println(ok)
	// Output: false
}

func ExampleCache_Clear() {
	c := cache.New[string, int](0)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Clear()
	fmt.Println(c.Count())
	// Output: 0
}

func ExampleCache_Count() {
	c := cache.New[string, int](0)
	c.Set("a", 1)
	c.Set("b", 2)
	fmt.Println(c.Count())
	// Output: 2
}

func ExampleCache_Has() {
	c := cache.New[string, int](0)
	c.Set("a", 1)
	fmt.Println(c.Has("a"))
	fmt.Println(c.Has("b"))
	// Output:
	// true
	// false
}
