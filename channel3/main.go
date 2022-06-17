package main

import (
	"fmt"
)

func fibonachi(ch chan int, mes chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-mes:
			fmt.Printf("Quit\n")
			return
		}
	}
}

func main() {
	c := make(chan int)
	m := make(chan bool)
	n := 10
	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-c)
		}
		m <- true
	}(n)
	fibonachi(c, m)
}
