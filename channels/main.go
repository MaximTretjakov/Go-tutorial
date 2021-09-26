package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("Executing goroutine")
	time.Sleep(1 * time.Second)
	c <- "Hello World"
	fmt.Println("Finished Executing goroutine")
}

func main() {
	fmt.Println("Hello from channels")

	values := make(chan string, 2)
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	value := <-values

	fmt.Println(value)

	time.Sleep(1 * time.Second)
}
