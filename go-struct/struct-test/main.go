package main

import (
	"fmt"

	"github.com/MaximTretjakov/Go-tutorial/go-struct/model"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p Person
	fmt.Println(p)

	p1 := Person{"Maxim", "Tretyakov", 40}
	fmt.Println(p1)

	p2 := Person{
		FirstName: "Maxim",
		LastName:  "Tretyakov",
		Age:       40,
	}
	fmt.Println(p2)

	p3 := Person{FirstName: "Maxim"}
	fmt.Println(p3)

	fmt.Println("FirstName", p3.FirstName)

	type Student struct {
		Number int
		Name   string
	}

	student := Student{80, "Alex"}
	t := &student

	fmt.Println(t)
	fmt.Println((*t).Name)
	fmt.Println(t.Number)

	type Employee struct {
		Id   int
		Name string
	}

	pEmp := new(Employee)
	pEmp.Id = 1
	pEmp.Name = "Stalin"
	fmt.Println(pEmp)

	c := model.Customer{
		Id:   1,
		Name: "Maxim Tretyakov",
	}
	// c.married:= true
	// a := model.address{}

	fmt.Println("Programmer : ", c)
}
