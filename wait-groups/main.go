package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished executing goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Hello from wait groups tutorial")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()
	fmt.Println("Finished executing my go program")
}
