package main

import (
	"fmt"
	"math/rand"
)

func uniqRandn(n int) []int {
	res, resMap := make([]int, 0, n), make(map[int]struct{}, n)
	for len(res) < n {
		val := rand.Int()
		if _, ok := resMap[val]; ok {
			continue
		}
		res = append(res, val)
		resMap[val] = struct{}{}
	}
	return res
}

func main() {
	fmt.Println(uniqRandn(10))
}
