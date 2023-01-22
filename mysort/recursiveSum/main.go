package main

import "fmt"

func sum(arr []int, n int) int {
	if n <= 0 {
		return 0
	}
	test := arr[n-1]
	fmt.Println(test)
	return (sum(arr, n-1) + arr[n-1])
}

func main() {
	arr := []int{1, 2, 7}
	fmt.Println(sum(arr, 3))
}
