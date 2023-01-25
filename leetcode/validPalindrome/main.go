/*
Задача Valid Palindrome решается с помошью паттрена two pointers
Это шаблон, в котором два указателя перебирают структуру данных в тандеме, пока один или оба указателя не достигнут определённого условия. 
Два указателя часто полезны при поиске пар в отсортированном массиве или связанном списке. 
Например, когда нужно сравнить каждый элемент массива с другими его элементами.
*/
package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		l, r := rune(s[left]), rune(s[right])
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
		} else if unicode.ToLower(l) == unicode.ToLower(r) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func main() {
	test1 := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(test1))
	test2 := "race a car"
	fmt.Println(isPalindrome(test2))
}
