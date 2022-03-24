package main

import (
	"fmt"
	"sync"
)

var (
	wg      sync.WaitGroup
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	wg.Add(1)

	// Alice
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Println("=", Balance())
	}()

	// Bob
	go Deposit(100)

	wg.Wait()
}
