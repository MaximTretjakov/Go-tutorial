package main

import "fmt"

func twoSum(nums []int, target int) []int {
	storage := make(map[int]int)
	for ind, val := range nums {
		x := target - val
		if val, ok := storage[x]; ok {
			return []int{val, ind}
		}
		storage[val] = ind
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 17))
}
