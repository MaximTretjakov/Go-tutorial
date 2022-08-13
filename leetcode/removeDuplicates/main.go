package main

import "fmt"

func removeDuplicates(nums []int) int {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1000 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums[j] = 1000
				counter++
			}
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
	fmt.Print(nums)
	return len(nums) - counter
}

func main() {
	// nums1 := []int{0, 0, 1, 1, 2}
	nums2 := []int{-100, -100, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums2)
	fmt.Printf("Unique: %d\n", k)
}
