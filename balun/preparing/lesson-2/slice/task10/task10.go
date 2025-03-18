package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Memory leaking example
	// data := findData()
	// _ = data

	// printAllocs()
	// runtime.GC()
	// printAllocs()

	// runtime.KeepAlive(data)

	// Prevent leaking by copy
	numbers := make([]int, 1<<30)
	pointer := findElement(numbers, 0)
	_ = pointer

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(pointer)
}

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func findData() []byte {
	data := make([]byte, 1<<30)

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			// return data[i : i+20]
			partData := make([]byte, 20)
			copy(partData, data)
			return partData
		}
	}

	return nil
}

func findElement(numbers []int, target int) *int {
	for i := range numbers {
		if numbers[i] == target {
			return &numbers[i]
		}
	}

	return nil
}
