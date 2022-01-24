package main

import "fmt"

func linearSearch(target int, mas []int) bool {
	for _, item := range mas {
		if item == target {
			return true
		}
	}

	return false
}

func main() {
	iterable := []int{1, 2, 3, 44, 0, 598, 12, 338, 123}
	fmt.Println(linearSearch(123, iterable))
}
