package main

import "fmt"

func main() {
	s := "test"
	fmt.Println(s[0])
	tmp := []byte(s)
	tmp[0] = 'R'
	fmt.Println(string(tmp))
}
