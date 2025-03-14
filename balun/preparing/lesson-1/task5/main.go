package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var number int32 = 0x12345678
	// pointer := unsafe.Pointer(&number)

	// fmt.Printf("0x")
	// for i := 0; i < 4; i++ {
	// 	byteValue := *(*int8)(unsafe.Add(pointer, i))
	// 	fmt.Printf("%x", byteValue)
	// }
	// fmt.Println()

	if isLittleEndian() {
		fmt.Println("Little endian")
		return
	}
	fmt.Println("Big endian")
}

func isLittleEndian() bool {
	var number int16 = 0x1000
	ptr := (*int8)(unsafe.Pointer(&number))
	return *ptr == 1
}

func isBigEndian() bool {
	return !isLittleEndian()
}
