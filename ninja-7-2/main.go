package main

import (
	"fmt"
)

type person struct {
	title string
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		title: "Mr.",
		first: "Mukul",
		last:  "Debnath",
		age:   32,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.title = "Unknown"
	p.first = "First"
	p.last = "Last"
	p.age = 42
}
