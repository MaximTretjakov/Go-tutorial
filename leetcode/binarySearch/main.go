package main

import "fmt"

func BinarySearch(m []int, target int) int {
	low := 0
	height := len(m) - 1

	for low <= height {
		mid := low + height
		guess := m[mid]
		if guess == target {
			return mid
		}
		if guess > target {
			height = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	mas := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(BinarySearch(mas, 4))
}
