/*
	Liskov substitution principle.
	Teacher и student "подтипы" human типа, и оба могут быть замещены human типом.
	В golang нет ООП но есть композиция, за счет этого мы имеем доступ к полям и методам вложенного типа.
*/

package main

import "fmt"

type person interface {
	getName() string
}

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	salary float64
}

type student struct {
	human
	marks float64
}

type printer struct{}

func (pr printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	h := human{name: "John"}

	s := student{
		human: human{name: "Bob"},
		marks: 90,
	}

	t := teacher{
		human:  human{name: "Prof Smith"},
		salary: 50000.0,
	}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
