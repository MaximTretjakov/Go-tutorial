package main

import (
	"fmt"
	"unsafe"
)

func main() {
	values := [...]int{100, 200, 300}

	for idx, value := range values {
		value += 50
		fmt.Println("#1:", unsafe.Pointer(&value), "#2:", unsafe.Pointer(&values[idx]))
	}

	fmt.Println("values:", values)

	iterationOverArray()
}

func iterationOverArray() {
	values := [...]int{100, 200, 300}

	// copy of array
	for _, v := range values {
		fmt.Println(v)
	}

	// not a copy
	for _, v := range &values {
		fmt.Println(v)
	}

	// not a copy
	for _, v := range values[:] {
		fmt.Println(v)
	}
}
