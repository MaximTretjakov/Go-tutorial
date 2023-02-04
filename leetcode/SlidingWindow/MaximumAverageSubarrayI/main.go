/*
Задача Maximum Average Subarray I должна решаться с помошью паттрена sliding window.
*/
package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	max := sum
	pointer := k

	for pointer < len(nums) {
		sum = sum - nums[pointer-k] + nums[pointer]
		if sum > max {
			max = sum
		}
		pointer++
	}

	return float64(max) / float64(k)
}

func main() {
	fmt.Printf("%v\n", findMaxAverage([]int{3, 3, 4, 3, 0}, 3))
	// fmt.Println(findMaxAverage([]int{}, 4))
	// fmt.Println(findMaxAverage([]int{100, 200, 300, 400}, 2))
}
