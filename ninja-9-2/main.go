package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hi, I am", p.first, p.last, "of age", p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Mukul",
		last:  "Debnath",
		age:   32,
	}
	//saySomething(p1)
	saySomething(&p1)
}
