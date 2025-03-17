package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4}

	fmt.Println("initial slice:", data)
	process1(data)
	fmt.Println("after process1:", data)
	process2(data)
	fmt.Println("after process2:", data)

	//slicing
	slice1()
}

func process1(data []int) {
	data[2] = 5
}

func process2(data []int) {
	data = append(data, 6)
	fmt.Println("len:", len(data), "cap:", cap(data))
	fmt.Println(data)
}

func slice1() {
	data := make([]int, 3, 6)

	fmt.Println("initial slice:", data)
	process(data)
	fmt.Println("after process:", data)
	fmt.Println("after process:", data[:4])
}

func process(data []int) {
	data = append(data, 55)
	// fmt.Println("in process:", data[:4])
}
