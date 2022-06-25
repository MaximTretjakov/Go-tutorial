package main

import "fmt"

func main() {
	const s = "Привет мир!"

	fmt.Printf("Len of s: %d\n", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println("For loop print bytes: \n", s[i])
	}

	fmt.Println()

	for i, v := range s {
		fmt.Printf("For loop: %#U starts at %d\n", v, i)
	}
}
