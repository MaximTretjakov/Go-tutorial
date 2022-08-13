package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	for i := range nums {
		if nums[i] == val {
			count++
			nums[i] = 1000
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}
	return len(nums) - count
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	k := removeElement(nums, 2)
	fmt.Println(k)
}
