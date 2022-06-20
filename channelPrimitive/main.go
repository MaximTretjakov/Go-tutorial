package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {
	fmt.Println("Begin work")
	time.Sleep(1 * time.Second)
	fmt.Println("End work")
	ch <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done
}
