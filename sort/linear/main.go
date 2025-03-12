package main

import (
	"fmt"
)

func Search(data []int, target int) bool {
	for _, item := range data {
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	data := []int{1, 2, 3, 4, 10, 113, 9, 99, 90}
	fmt.Println(Search(data, 5))
}
