package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Lenght, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Lenght * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Lenght + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2*math.Pi + c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

func main() {
	var s Shape = Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())
}
