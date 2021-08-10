package main

import (
	"fmt"
)

func printEmployee(data map[string]int, id string) {
	if val, ok := data[id]; ok {
		fmt.Printf("value = %d, ok = %v\n", val, ok)
	} else {
		fmt.Printf("id = %s, not found\n", id)
	}
}

func isEmployeeExist(data map[string]int, id string) bool {
	_, ok := data[id]
	return ok
}

func main() {
	var my_first_map map[int]string

	fmt.Println(my_first_map)

	if my_first_map == nil {
		fmt.Printf("My %v is nil\n", my_first_map)
	}

	// get error
	// my_first_map[1] = "test"

	// init map by make()
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Print("map if nil")
	} else {
		fmt.Print("map initilized\n")
	}

	m["test"] = 123
	m["qwerty"] = 222
	fmt.Printf("%v\n", m)

	// init map by literal
	var x = map[int]int{
		1: 123,
		2: 321,
		3: 444,
	}

	fmt.Println(x)

	var y = x[3]
	fmt.Printf("get value by key : %v\n", y)
	fmt.Printf("key does not exist : %v\n", x[100])

	foo, ok := x[1]
	if ok {
		fmt.Printf("key 1 exist %v\n", foo)
	} else {
		fmt.Println("key does not exist")
	}

	boo, ok := x[1000]
	if ok {
		fmt.Printf("key 1000 exist %v\n", boo)
	} else {
		fmt.Println("key 1000 does not exist")
	}

	// functions
	printEmployee(m, "test")
	printEmployee(m, "None")
	if isEmployeeExist(m, "test") {
		fmt.Println("EmployeeId bu found")
	}

	// delete key by delete()
	delete(m, "test")
	fmt.Println(m)

	for tt, yy := range x {
		fmt.Printf("Iterate over map: key = %d value = %d\n", tt, yy)
	}
}
