package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(palindrome int) bool {
	str := strconv.Itoa(palindrome)
	first := 0
	last := len(str) - 1

	for first <= last {
		if len(str) == 0 {
			return true
		}

		if str[first] != str[last] {
			return false
		}

		first++
		last--
	}

	return true
}

func main() {
	palindrome := 1221
	fmt.Println(isPalindrome(palindrome))
}
