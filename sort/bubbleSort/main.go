package main

import (
	"math/rand"
	"time"
)

func generateSlice(count int) []int {
	slice := make([]int, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func bubbleSort(slice []int) {
	n := len(slice)
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i+1], slice[i] = slice[i], slice[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func main() {
	sl := generateSlice(20)
	bubbleSort(sl)
}
