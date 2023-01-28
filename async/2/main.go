package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	store := make(map[int]int, writes)

	wg.Add(writes)

	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			store[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(store)
}
