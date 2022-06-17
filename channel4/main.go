package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	ch <- sum
}

func main() {
	mas := []int{17, 12, 18, 9, 24, 42, 64, 12, 68, 82, 57, 32, 9, 2, 12, 82, 52, 34, 92, 36, 12, 444, 45, 61, 90}

	fmt.Printf("Len: %d\n", len(mas))

	ch := make(chan int)

	for i := 0; i < len(mas); i = i + 5 {
		go sum(mas[i:i+5], ch)
	}

	output := make([]int, 5)

	for j := 0; j < 4; j++ {
		output[j] = <-ch
	}

	close(ch)

	fmt.Println(output)
}
