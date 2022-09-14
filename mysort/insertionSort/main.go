package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		s = append(s, rand.Intn(999)-rand.Intn(999))
	}
	return s
}

func insertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		j := i
		for j > 0 {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
			j -= 1
		}
	}
}

func main() {
	slice := generateSlice(20)
	fmt.Println("Unsorted slice ", slice)
	insertionSort(slice)
	fmt.Println("Sorted slice ", slice)
}
