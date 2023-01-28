package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000
	doubles := make([]int, 0, capacity)
	uniqeIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				alreadyStored[doubles[i]] = struct{}{}
				uniqeIDs <- doubles[i]
			}
		}()
	}

	close(uniqeIDs)
	wg.Done()
	for c := range uniqeIDs {
		fmt.Println(c)
	}

	fmt.Printf("len of ids: %d\n", len(uniqeIDs))
	fmt.Println(uniqeIDs)
}
