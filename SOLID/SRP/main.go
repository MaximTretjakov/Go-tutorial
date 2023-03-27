/*
	Плохой пример.
	Метод area производит вычисление и вывод результатов в консоль, что нарушает принцип SRP.
	В будущем возможно изменение форматов вывода (разные сообщения, форматы даннных JSON, XML и тд...)
	Метод area должен только вычислять и возвращать результат, а другие методы должны только отвечать за вывод информации.

	Метод goodArea и printArea демонстрируют правильный подход.
*/

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

// Bad implementation
func (c circle) area() {
	fmt.Println("Area : ", math.Pi*c.radius*c.radius)
}

// Bad implementation
func (s square) area() {
	fmt.Println("Area : ", s.length*s.length)
}

/*
	Хорошая реализация.
	Метод goodArea и printArea демонстрируют правильный подход.
*/

// Good implementation
func (c circle) goodArea() float64 {
	return math.Pi * c.radius * c.radius
}

// Good implementation
func (s square) goodArea() float64 {
	return s.length * s.length
}

func printArea(area float64) {
	fmt.Println("Good area : ", area)
}

func main() {
	// c := circle{radius: 1}
	// c.area()

	// s := square{length: 2}
	// s.area()

	c := circle{radius: 1}
	printArea(c.goodArea())

	s := square{length: 2}
	printArea(s.goodArea())
}
