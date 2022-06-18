package main

import "fmt"

func even(ch <-chan int) {
	for v := range ch {
		fmt.Printf("Even: %d\n", v)
	}
}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Printf("Odd: %d\n", v)
	}
}

func main() {
	mas := []int{91, 42, 23, 14, 15, 76, 87, 28, 19, 95}
	chEven := make(chan int)
	chOdd := make(chan int)

	go even(chEven)
	go odd(chOdd)

	for _, v := range mas {
		if v%2 != 0 {
			chOdd <- v
		} else {
			chEven <- v
		}
	}

}
