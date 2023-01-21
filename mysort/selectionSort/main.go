package main

import (
	"fmt"
)

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	arr := []int{145, 12, 34, 200, 0, 1, 15}
	fmt.Println(selectionSort(arr))
}
