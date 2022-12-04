package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	counter := 0
	for ind, val := range nums1 {
		if val == 0 {
			if counter == len(nums2){
				break
			}
			nums1[ind] = nums2[counter]
			counter++
		}
	}

	lenght := len(nums1)
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < len(nums1)-1; i++ {
			if nums1[i] > nums1[i+1] {
				nums1[i+1], nums1[i] = nums1[i], nums1[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		lenght -= 1
	}
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m := 6
	nums2 := []int{1, 2, 2}
	n := 3
	// output [1,2,2,3,5,6]
	merge(nums1, m, nums2, n)
}
