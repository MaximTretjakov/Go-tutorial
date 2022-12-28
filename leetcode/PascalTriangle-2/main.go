package main

import "fmt"

func getRow(rowIndex int) []int {
    result:=make([][]int, 0)
    for i:=1; i<=rowIndex+1; i++{
        if i == 1{
			result = append(result, []int{1})
			continue
		}
		if i == 2{
			result = append(result, []int{1, 1})
			continue
		}
		j:=i-2
		tmp:=[]int{1}
		for k:=0; k<j; k++{
			tmp = append(tmp, result[i-2][k]+result[i-2][k+1])
		}
		tmp = append(tmp, 1)
		result = append(result, tmp)
    }
	return result[rowIndex]
}

func main(){
	fmt.Println(getRow(3))
}