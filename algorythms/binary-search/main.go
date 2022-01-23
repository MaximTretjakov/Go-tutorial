package main

import (
	"fmt"
	"sort"
)

func binarySearch(target int, haystack []int) bool {
	begin := 0
	end := len(haystack) - 1

	for begin <= end {
		median := (begin + end) / 2

		if haystack[median] < target {
			begin = median + 1
		} else {
			end = median - 1
		}
	}

	if begin == len(haystack) || haystack[begin] != target {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 3, 89, 101, 4, 57, 22, 41, 222}
	fmt.Printf("Before sorting: %d\n", items)
	sort.Ints(items)
	fmt.Printf("After sorting: %d\n", items)
	fmt.Println(binarySearch(101, items))
}
