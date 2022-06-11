package main

import "fmt"

func reverse(sw []int) []int {
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		sw[a], sw[b] = sw[b], sw[a]
		fmt.Printf("\nA: %d B: %d", a, b)
		fmt.Printf("\nSlice len: %d", len(sw))
	}

	return sw
}

func main() {
	sw := []int{3, 2, 1}
	fmt.Print(reverse(sw))
}
