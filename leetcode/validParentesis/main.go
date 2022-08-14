package main

import (
	"fmt"
)

// func isValid(s string) bool {
// 	match := false
// 	if len(s)%2 == 0 {
// 		for i := range s {
// 			if match {
// 				match = false
// 				continue
// 			}
// 			if s[i] == '(' {
// 				if s[i+1] == ')' {
// 					match = true
// 				} else {
// 					return false
// 				}
// 			} else if s[i] == '[' {
// 				if s[i+1] == ']' {
// 					match = true
// 				} else {
// 					return false
// 				}
// 			} else if s[i] == '{' {
// 				if s[i+1] == '}' {
// 					match = true
// 				} else {
// 					return false
// 				}
// 			} else {
// 				return false
// 			}
// 		}
// 	} else {
// 		return false
// 	}
// 	return true
// }

func lenTwo(arr string, open rune, closed rune) bool {
	if rune(arr[0]) == open {
		if rune(arr[1]) == closed {
			return true
		}
	}
	return false
}

func searchSecond(arr string, ch rune) bool {
	// s2 := "{[]}"
	// s3 := "()[]{}"
	counter := 0
	for i := range arr {
		if ch == rune(arr[i+1]) {
			counter++
		}
	}
	return false
}

func isValid(s string) bool {
	if len(s)%2 == 0 {
		if len(s) == 2 {
			if lenTwo(s, '(', ')') {
				return true
			}
			if lenTwo(s, '[', ']') {
				return true
			}
			if lenTwo(s, '{', '}') {
				return true
			} else {
				return false
			}
		} else {
			for i := range s {
				if s[i] == '(' {
					if searchSecond(s, '(') {
						return true
					} else {
						return false
					}
				} else if s[i] == '{' {
					if searchSecond(s, '{') {
						return true
					} else {
						return false
					}
				} else if s[i] == '[' {
					if searchSecond(s, '[') {
						return true
					} else {
						return false
					}
				}

			}
		}
	} else {
		return false
	}
	return true
}

func main() {
	// s1 := "([)]"
	s2 := "{[]}"
	// s3 := "()[]{}"
	// s4 := "()"
	// s5 := "("
	fmt.Println(isValid(s2))
}
