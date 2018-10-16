package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "of age", p.age)
}

func main() {
	p1 := person{
		first: "Mukul",
		last:  "Debnath",
		age:   32,
	}
	p1.speak()
}
