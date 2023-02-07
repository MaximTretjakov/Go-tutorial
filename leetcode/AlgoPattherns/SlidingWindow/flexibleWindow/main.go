/*
Sliding window approach to maximum lenght range
*/
package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubstringlenght(ar []int, k int) int {
	i := 0
	j := 0
	sum := 0
	maxLen := 0

	for j < len(ar) {
		sum += ar[j]
		if sum < k {
			j++
		} else if sum == k {
			maxLen = max(maxLen, j-i+1)
			j++
		} else if sum > k {
			for sum > k {
				sum -= ar[i]
				i++
			}
			if sum == k {
				maxLen = max(maxLen, j-i+1)
			}
			j++
		}
	}

	return maxLen
}

func main() {
	m1 := []int{1, 2, 100, 32, 44, 4, 2, 1, 7}
	sum := 10
	fmt.Println(maxSubstringlenght(m1, sum))
}
