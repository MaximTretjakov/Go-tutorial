package main

import (
	"fmt"
	"sync"
)

func main() {
	ops := 1000
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	store := make(map[int]int, ops)

	wg.Add(ops)
	for i := 0; i < ops; i++ {
		i := i
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			store[i] = i
		}()
	}

	wg.Add(ops)
	for i := 0; i < ops; i++ {
		i := i
		go func() {
			defer wg.Done()
			defer mu.RUnlock()
			mu.RLock()
			_, _ = store[i]
		}()
	}

	wg.Wait()
	fmt.Println(store)
}
