package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, num := range nums {
		if value, ok := m[num]; ok {
			if absInt(i, value) <= k {
				return true
			} else {
				m[num] = i
			}
		} else {
			m[num] = i
		}
	}

	return false
}

func absInt(a, b int) int {
	if a-b < 0 {
		return b - a
	}

	return a - b
}

func main() {
	nums1 := []int{1, 2, 3, 1}
	k1 := 3
	fmt.Println(containsNearbyDuplicate(nums1, k1)) // true

	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	fmt.Println(containsNearbyDuplicate(nums2, k2)) // true

	nums3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 2
	fmt.Println(containsNearbyDuplicate(nums3, k3)) // false

	nums4 := []int{-1, -1}
	k4 := 1
	fmt.Println(containsNearbyDuplicate(nums4, k4)) // true

	nums5 := []int{1, 2, 1}
	k5 := 2
	fmt.Println(containsNearbyDuplicate(nums5, k5)) // true
}
