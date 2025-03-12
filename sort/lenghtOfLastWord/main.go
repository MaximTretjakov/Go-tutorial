package main

import "fmt"

func lengthOfLastWord(s string) int {
	wordLenght := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			wordLenght++
		} else {
			if wordLenght != 0 {
				break
			}
		}
	}
	return wordLenght
}

func main() {
	s := " Hel   lo Wor 1            "
	fmt.Println(lengthOfLastWord(s))
}
