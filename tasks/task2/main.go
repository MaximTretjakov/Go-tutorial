package main

import "fmt"

func twoNumberSum(array []int, target int) []int {
	matches := make(map[int]struct{})

	for _, i := range array {
		diff := target - i
		if _, ok := matches[i]; ok {
			return []int{i, diff}
		} else {
			matches[diff] = struct{}{}
		}
	}

	return []int{}
}

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	fmt.Println(twoNumberSum(array, 10))
}
