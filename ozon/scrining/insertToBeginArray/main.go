package main

import "fmt"

func remove(in []int) []int {
	res := make([]int, 0, len(in))

	for i := range in {
		if in[i] != 0 {
			res = append(res, in[i])
		}
	}

	return res
}

func main() {
	s3 := []int{1, 0, 0, 2}
	fmt.Println(remove(s3))
}
