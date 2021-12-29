package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
	getName() string
}

type Square struct {
	sideLength float64
	name       string
}

type Circle struct {
	radius float64
	name   string
}

func (c Circle) getName() string {
	return c.name
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius * math.Pi
}

func (s Square) getName() string {
	return s.name
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s Shape) {
	fmt.Println("Area of", s.getName(), "is", s.getArea())
}

func getName(s Shape) string {
	return s.getName()
}

func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {
	s := Square{
		sideLength: 5.5,
		name:       "Square",
	}

	c := Circle{
		radius: 5.5,
		name:   "Circle",
	}

	printArea(s)
	printArea(c)
}
