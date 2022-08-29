package main

import "fmt"

func zip(s ...[]int) [][]int {
	if len(s) == 0 {
		return [][]int{}
	}

	minLen := len(s[0])

	for i := 1; i < len(s); i++ {
		if len(s[i]) < minLen {
			minLen = len(s[i])
		}
	}

	res := make([][]int, 0, minLen)

	for i := 0; i < minLen; i++ {
		x := make([]int, 0, len(s))
		for j := 0; j < len(s); j++ {
			x = append(x, s[j][i])
		}
		res = append(res, x)
	}

	return res
}

func main() {
	s1, s2, s3 := []int{1, 2, 3}, []int{4, 5, 6, 7}, []int{11, 12, 13, 14}
	fmt.Println(zip(s1, s2, s3))
}
