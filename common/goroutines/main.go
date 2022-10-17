package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a(ch chan string) {
	for i := 0; i < 50; i++ {
		ch <- "a"
		time.Sleep(1 * time.Second)
	}
	defer wg.Done()
}

func b(ch chan string) {
	for i := 0; i < 50; i++ {
		ch <- "b"
		time.Sleep(2 * time.Second)
	}
	defer wg.Done()
}

func main() {
	c := make(chan string)
	wg.Add(1)
	go a(c)
	wg.Add(1)
	go b(c)

	go func() {
		for m := range c {
			fmt.Println(m)
		}
	}()

	wg.Wait()
}
