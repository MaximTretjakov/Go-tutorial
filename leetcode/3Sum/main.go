/*
Задача Reverse Words in a String должна решаться с помошью паттрена two pointers, но я решить не смог.
*/
package main

import (
	"fmt"
	"unicode"
)

// func reverseWords(s string) string {
// 	store := [][]rune{}
// 	symbolLeft := 0
// 	symbolRight := 0
// 	left, right := 0, len(s)-1
// 	tmp := []rune(s)

// 	for left < right {
// 		l, r := rune(s[left]), rune(s[right])
// 		nextLeft, nextRight := rune(s[left+1]), rune(s[right-1])
// 		if unicode.IsLetter(l) || unicode.IsDigit(l) {
// 			symbolLeft++
// 			left++
// 			if !unicode.IsLetter(nextLeft) && !unicode.IsDigit(nextLeft) {
// 				store = append(store, tmp[left-symbolLeft:left])
// 				symbolLeft = 0
// 			}
// 		} else if unicode.IsLetter(r) || unicode.IsDigit(r) {
// 			symbolRight++
// 			if !unicode.IsLetter(nextRight) && !unicode.IsDigit(nextRight) {
// 				store = append(store, tmp[right:right+symbolRight])
// 				symbolRight = 0
// 			}
// 			right--
// 		} else {
// 			left++
// 			right--
// 		}
// 	}
// 	return ""
// }

func reverseWords(s string) string {
	store := ""
	symbolRight := 0
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			symbolRight++
		} else {
			word := s[i+1 : i+1+symbolRight]
			if len(word) != 0 {
				store += word + " "
				symbolRight = 0
			}
		}
		if i == 0 {
			word := s[i : i+symbolRight]
			if len(word) != 0 {
				store += word
				symbolRight = 0
			}
		}
	}
	last:=rune(store[len(store)-1])
	if !unicode.IsLetter(last) && !unicode.IsDigit(last){
		return store[:len(store)-1]
	}
	return store
}

func main() {
	// s1 := "the sky is blue"
	// fmt.Println(reverseWords(s1))
	s2 := "  hello world  "
	fmt.Println(reverseWords(s2))
	// s3 := "a good   example"
	// fmt.Println(reverseWords(s3))
}
