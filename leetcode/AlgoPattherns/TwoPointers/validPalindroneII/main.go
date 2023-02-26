package main

import "fmt"

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	test1 := "aba"
	test2 := "abca"
	test3 := "abc"
	test4 := "abfda"
	test5 := "ababza"
	test6 := "deeee"
	fmt.Println(validPalindrome(test1))
	fmt.Println(validPalindrome(test2))
	fmt.Println(validPalindrome(test3))
	fmt.Println(validPalindrome(test4))
	fmt.Println(validPalindrome(test5))
	fmt.Println(validPalindrome(test6))
}
