package main

import (
	"fmt"
	"math/rand"
)

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{124, 17, 0, 33, 55, 610, 111, 359, 1}
	fmt.Println(quicksort(arr))
}
