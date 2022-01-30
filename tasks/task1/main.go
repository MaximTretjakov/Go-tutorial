package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for i := 0; i < len(str); i++ {
		if left >= right {
			return true
		}

		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	s := "abcdcba"

	fmt.Print(isPalindrome(s))
}
