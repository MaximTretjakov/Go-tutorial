package main

import "fmt"

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotal(triangle [][]int) int {
	rowSize := len(triangle)
	for row := rowSize - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}
