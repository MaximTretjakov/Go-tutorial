/*
	Interface segregation principle.
	Плохой пример.
	Тут нарушен принцип, тип square должен реализовать метод volume хоть и не нуждается в нем иначе он не реализует интерфейс.
*/

package main

import "fmt"

// type shape interface {
// 	area() float64
// 	volume() float64
// }

// func areaSum(shapes ...shape) float64 {
// 	var sum float64

// 	for _, shape := range shapes {
// 		sum += shape.area()
// 	}

// 	return sum
// }

// func areaVolumeSum(shapes ...shape) float64 {
// 	var sum float64

// 	for _, shape := range shapes {
// 		sum += shape.volume() + shape.area()
// 	}

// 	return sum
// }

// type square struct {
// 	length float64
// }

// func (s square) area() float64 {
// 	return s.length * s.length
// }

// func (s square) volume() float64 {
// 	return 0
// }

// type cube struct {
// 	length float64
// }

// func (c cube) area() float64 {
// 	return 6.0 * c.length * c.length
// }

// func (c cube) volume() float64 {
// 	return c.length * c.length * c.length
// }

/*
	Правильная реализация.
	Нужно разделить интерфейсы.
	Для каждого типа определить свой интерфейс и реализовать его.
	Исключить принудительную реализацию методов интерфейса которые не нужны конкретному классу.
*/

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

func areaSum(shapes ...shape) float64 {
	var sum float64

	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func areaVolumeSum(shapes ...object) float64 {
	var sum float64

	for _, shape := range shapes {
		sum += shape.volume() + shape.area()
	}

	return sum
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type cube struct {
	length float64
}

func (c cube) area() float64 {
	return 6.0 * c.length * c.length
}

func (c cube) volume() float64 {
	return c.length * c.length * c.length
}

func main() {
	// s := square{length: 2.0}
	// c := cube{length: 23.0}

	// fmt.Println(areaSum(s, c))
	// fmt.Println(areaVolumeSum(s, c))

	s := square{length: 1.0}
	c := cube{length: 2.0}

	fmt.Println(areaSum(s))
	fmt.Println(areaVolumeSum(c))
}
