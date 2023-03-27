/*
	Плохой пример.
	При добавлении нового case в метод sumArea нарушается принцип OCP.
	Объект должен быть закрыт для изменения и открыт для расширения.
*/

package main

import (
	"fmt"
	"math"
)

// type circle struct {
// 	radius float64
// }

// type square struct {
// 	lenght float64
// }

// type calculator struct{}

// func (c calculator) sumArea(shapes ...interface{}) float64 {
// 	var sum float64

// 	for _, shape := range shapes {
// 		switch shape.(type) {
// 		case circle:
// 			r := shape.(circle).radius
// 			sum += math.Pi * r * r
// 		case square:
// 			l := shape.(square).lenght
// 			sum += l * l
// 		}
// 	}

// 	return sum
// }

/*
	Хороший пример.
	Используя абстракцию (интерфейс) мы не изменяем код sumArea, а расширяем за счет нового типа реализующего метод интерфейса shape.
	Реализовав интерфейс мы можем прокидывать тип в метод с интерфейсным аргументом (неявное приведение к типу интерфейса),
	расширяя его таким образом.
	Этот подход способствует исключению ошибок, т.к мы не меняем код.
*/

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	lenght float64
}

func (s square) area() float64 {
	return s.lenght * s.lenght
}

type calculator struct{}

func (c calculator) sumArea(shapes ...shape) float64 {
	var sum float64

	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func main() {
	c := circle{radius: 2}
	s := square{lenght: 4}

	calc := calculator{}
	sum := calc.sumArea(c, s)

	fmt.Println("Area sum: ", sum)
}
