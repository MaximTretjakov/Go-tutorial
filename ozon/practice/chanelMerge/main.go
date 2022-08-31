package main

import (
	"fmt"
	"sync"
)

func merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(ch))

	for _, c := range ch {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1, ch2 := make(chan int, 1), make(chan int, 1)

	ch1 <- 5
	close(ch1)
	ch2 <- 4
	close(ch2)

	for c := range merge(ch1, ch2) {
		fmt.Println(c)
	}
}
