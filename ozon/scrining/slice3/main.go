package main

import (
	"fmt"
	"time"
)

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}

func main() {
	start := time.Now()
	_, _ = <-worker(), <-worker()
	fmt.Println(int(time.Since(start).Seconds()))
}
