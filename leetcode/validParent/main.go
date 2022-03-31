package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func isValid(s string) bool {
	st := stack.New()

	for _, v := range s {
		if v == '(' {
			st.Push(')')
		} else if v == '[' {
			st.Push(']')
		} else if v == '{' {
			st.Push('}')
		} else {
			if v != st.Pop() {
				break
			}
		}
	}
	return st.Len() == 0
}

func main() {
	// s1 := "()"
	s2 := "()[]{}"
	// s3 := "()["
	// s4 := "(["

	fmt.Println(isValid(s2))
}
