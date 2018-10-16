package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func intfArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		length: 2,
	}
	c := circle{
		radius: 2,
	}
	intfArea(s)
	intfArea(c)
	fmt.Println(s.area())
	fmt.Println(c.area())
}
