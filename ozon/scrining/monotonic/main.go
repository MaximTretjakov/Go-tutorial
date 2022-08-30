package main

import "fmt"

func isMonotonic(in []int) bool {
	isDown, isUp := true, true
	for i := 1; i < len(in); i++ {
		isDown = isDown && in[i-1] >= in[i]
		isUp = isUp && in[i-1] <= in[i]
	}
	return isUp || isDown
}

func main() {
	s1 := []int{23, 5, 23}
	fmt.Println(isMonotonic(s1))
}
