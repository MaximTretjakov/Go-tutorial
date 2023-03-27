/*
	Правильная реализация.
	Сущности (функции\структуры поля\методы) должны зависить от абстракций (интерфейсов) не от конкретной реализации.
*/

package main

import "fmt"

type greet interface {
	Greet() string
}

type french struct{}

func (f french) Greet() string {
	return "Bonjour"
}

type english struct{}

func (e english) Greet() string {
	return "Hello"
}

func greeter(g greet) {
	fmt.Println(g.Greet())
}

func main() {
	e := english{}
	greeter(e)

	f := french{}
	greeter(f)
}
