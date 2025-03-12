package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "lo"
	fmt.Println(strStr(haystack, needle))
}
