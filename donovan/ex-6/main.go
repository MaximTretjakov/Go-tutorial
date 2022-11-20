package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	chRND := make(chan int, 10)
	chRNDQuad := make(chan int, 10)
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			rndNum := rand.Intn(10-1) + 1
			fmt.Printf("Generate random numbers %d\n", rndNum)
			chRND <- rndNum
		}()
	}

	go func() {
		wg2.Wait()
		close(chRND)
	}()

	for d := range chRND {
		wg1.Add(1)
		go func(d int) {
			defer wg1.Done()
			result := d * 2
			fmt.Printf("Build quads %d\n", result)
			chRNDQuad <- result
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
