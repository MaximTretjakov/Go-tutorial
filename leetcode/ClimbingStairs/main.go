package main

import "fmt"

func ClimbStairs(n int) int {
	res := 0
	last := 0
	secondLast := 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			res = 1
		} else if i == 2 {
			res = 2
		} else {
			res = last + secondLast
		}
		secondLast = last
		last = res
		fmt.Printf("secondLast = %d\n", secondLast)
		fmt.Printf("last = %d\n", last)
		fmt.Printf("result = %d\n", res)
	}
	return res
}

func main() {
	fmt.Printf("Climb Stairs: %d\n", ClimbStairs(11))
}
