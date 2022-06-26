package main

import "fmt"

func Summ(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("%d + %d = %d\n", 2, 2, Summ(2, 2))
}
