package main

import "fmt"

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert(nums, target))
}
