package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 1; i < numRows+1; i++ {
		if i == 1 {
			result = append(result, []int{1})
			continue
		}
		if i == 2 {
			result = append(result, []int{1, 1})
			continue
		}
		j := i - 2
		previous := result[j]
		tmp := []int{1}
		for k := 0; k < j; k++ {
			previousSum := previous[k] + previous[k+1]
			tmp = append(tmp, previousSum)
		}
		tmp = append(tmp, 1)
		result = append(result, tmp)
	}
	return result
}

func main() {
	fmt.Println(generate(1))
}
