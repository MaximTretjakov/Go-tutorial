package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append(digits, 1)
	tmp := digits[len(digits)-1]
	digits[len(digits)-1] = digits[0]
	digits[0] = tmp
	return digits
}

func main() {
	digits1 := []int{9, 9}
	fmt.Println(plusOne(digits1))
}
