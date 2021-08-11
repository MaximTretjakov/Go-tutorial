package main

import (
	"fmt"
)

func main() {
	var test1 = []int{1, 2, 3}
	fmt.Printf("slice from literal : %d\n", test1)

	var test2 = [5]int{1, 2, 3, 4, 5}
	var test2_slice []int = test2[1:4]
	fmt.Printf("Array : %d\n", test2)
	fmt.Printf("Slice : %d\n", test2_slice)

	test3 := [5]string{"C", "C++", "Python", "Ruby", "GO"}
	fmt.Printf("Array of lanuages : %v\n", test3)

	test4_slice := test3[1:3]
	fmt.Printf("Slice 1:3 %v\n", test4_slice)

	test5_slice := test3[:]
	fmt.Printf("Slice all %v\n", test5_slice)

	test6_slice := test3[:3]
	fmt.Printf("Slice :3 %v\n", test6_slice)

	test7_slice := test3[2:]
	fmt.Printf("Slice 2: %v\n", test7_slice)

	test8_slice := make([]int, 5, 10)
	fmt.Printf("Slice by make() slice=%d lenght=%d capacity=%d", test8_slice, len(test8_slice), cap(test8_slice))
}
