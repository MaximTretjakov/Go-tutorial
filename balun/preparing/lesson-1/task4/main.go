package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// go run -gcflags=-d=checkptr main.go
	// x := new(int)
	// y := new(int)
	// z := new(int)

	// ptrX := unsafe.Pointer(x)
	// ptrY := unsafe.Pointer(y)
	// addressZ := uintptr(unsafe.Pointer(z))

	// // Арифметика указателей
	// _ = addressZ + 2
	// _ = addressZ - 2

	// runtime.GC()

	// *(*int)(ptrX) = 100
	// *(*int)(ptrY) = 200
	// *(*int)(unsafe.Pointer(addressZ)) = 300 // dangerous

	// task4
	var array [10]int
	address1 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#1 array address:", address1)

	allocation(10000000)

	address2 := (uintptr)(unsafe.Pointer(&array))
	fmt.Println("#2 array address:", address2)
	fmt.Println("#1 array address:", address1)
}

//go:noinline
func allocation(index int) byte {
	var data [200000000]byte
	return data[index]
}
