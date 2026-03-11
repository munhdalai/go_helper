package helpers_test

import (
	"fmt"

	"github.com/munhdalai/go_helper/helpers"
)

func ExampleContains() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(helpers.Contains(nums, 3))
	fmt.Println(helpers.Contains(nums, 6))
	// Output:
	// true
	// false
}

func ExampleContains_string() {
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println(helpers.Contains(fruits, "banana"))
	// Output: true
}

func ExampleMap() {
	nums := []int{1, 2, 3, 4}
	doubled := helpers.Map(nums, func(n int) int {
		return n * 2
	})
	fmt.Println(doubled)
	// Output: [2 4 6 8]
}

func ExampleMap_toString() {
	nums := []int{1, 2, 3}
	strs := helpers.Map(nums, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println(strs)
	// Output: [#1 #2 #3]
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := helpers.Filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evens)
	// Output: [2 4 6]
}

func ExampleReduce() {
	nums := []int{1, 2, 3, 4, 5}
	sum := helpers.Reduce(nums, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println(sum)
	// Output: 15
}

func ExampleReduce_concat() {
	words := []string{"Hello", " ", "World"}
	result := helpers.Reduce(words, "", func(acc string, s string) string {
		return acc + s
	})
	fmt.Println(result)
	// Output: Hello World
}

func ExampleChunk() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	chunks := helpers.Chunk(nums, 3)
	fmt.Println(chunks)
	// Output: [[1 2 3] [4 5 6] [7]]
}

func ExampleUnique() {
	nums := []int{1, 2, 2, 3, 3, 3, 4}
	result := helpers.Unique(nums)
	fmt.Println(result)
	// Output: [1 2 3 4]
}

func ExampleUnique_string() {
	colors := []string{"red", "blue", "red", "green", "blue"}
	result := helpers.Unique(colors)
	fmt.Println(result)
	// Output: [red blue green]
}

func ExamplePaginateSlice() {
	items := []string{"a", "b", "c", "d", "e", "f", "g"}

	page1 := helpers.PaginateSlice(items, 1, 3)
	fmt.Println(page1)

	page2 := helpers.PaginateSlice(items, 2, 3)
	fmt.Println(page2)

	page3 := helpers.PaginateSlice(items, 3, 3)
	fmt.Println(page3)
	// Output:
	// [a b c]
	// [d e f]
	// [g]
}
