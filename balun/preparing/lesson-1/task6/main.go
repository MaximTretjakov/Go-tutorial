package main

import "fmt"

func main() {
	var val1 int32 = 100
	ptr := &val1

	fmt.Println("value:", *ptr)
	fmt.Println("address:", ptr)

	process(&ptr)

	fmt.Println("value:", *ptr)
	fmt.Println("address:", ptr)
}

func process(temp **int32) {
	var val2 int32 = 200
	*temp = &val2
}
