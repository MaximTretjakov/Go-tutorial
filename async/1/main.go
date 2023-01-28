package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 20
	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}
	wg.Wait()
}
