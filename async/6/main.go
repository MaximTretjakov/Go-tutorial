package main

import (
	"fmt"
	"sync"
)

func merge[T any](channels ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{}
	for _, singleChan := range channels {
		singleChan := singleChan
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range singleChan {
				result <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4

	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}
