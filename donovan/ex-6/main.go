package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	chRND := make(chan int64, 10)
	chRNDQuad := make(chan int64, 10)
	var wg1 sync.WaitGroup

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		rndNum := rand.Int63()
		fmt.Println("Generate random numbers")
		chRND <- rndNum
	}
	close(chRND)

	for d := range chRND {
		wg1.Add(1)
		go func(d int64) {
			defer wg1.Done()
			fmt.Println("Build quads")
			chRNDQuad <- d * 2
		}(d)
	}

	go func() {
		wg1.Wait()
		close(chRNDQuad)
	}()

	for quad := range chRNDQuad {
		fmt.Printf("Read quads\t%d\n", quad)
	}
}
