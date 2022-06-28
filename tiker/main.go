package main

import (
	"fmt"
	"time"
)

func main() {
	tiker := time.NewTicker(500 * time.Microsecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-tiker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Microsecond)
	tiker.Stop()
	done <- true
	fmt.Println("Ticker stopped!")
}
