package main

import "fmt"

func zip(s1 []int, s2 []int) [][]int {
	res := [][]int{}
	r := []int{}
	lenght := 0

	if len(s1) < len(s2) {
		lenght = len(s1)
	} else {
		lenght = len(s2)
	}

	for i := 0; i < lenght; i++ {
		r = append(r, s1[i], s2[i])
		res = append(res, r)
		r = nil
	}
	return res
}

func ozonZip(s1 []int, s2 []int) [][]int {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	res := make([][]int, 0, minLen)
	for i := 0; i < minLen; i++ {
		res = append(res, []int{s1[i], s2[i]})
	}
	return res
}

func main() {
	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6, 7, 8}
	// fmt.Println(zip(s1, s2))
	fmt.Println(ozonZip(s1, s2))
}
