package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int) {
	min := 1
	max := 5

	delay := rand.Intn(max-min+1) + min

	fmt.Printf("Worker start: %d\n", id)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("Worker end: %d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
