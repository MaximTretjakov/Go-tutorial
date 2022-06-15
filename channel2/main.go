package main

import "fmt"

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func receiver(c <-chan int) {
	for v := range c {
		fmt.Printf("Receiver: %d\n", v)
	}
}

func main() {
	c := generator()
	receiver(c)
}
