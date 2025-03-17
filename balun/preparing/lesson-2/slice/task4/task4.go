package main

import "fmt"

// x - [ "A", "Z", "Z", "Z", "Z" ]
// x - [ "A", "M", "C" ]
// i = 0
// i = 1

func main() {
	var x = []string{"A", "M", "C"}
	fmt.Printf("Array len main: %d cap: %d\n", len(x), cap(x))

	for i, s := range x {
		fmt.Printf("Array len range: %d cap: %d\n", len(x), cap(x))
		println(i, s)
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
