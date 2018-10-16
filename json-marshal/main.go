package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First      string
	Last       string
	Age        int
	favFlavors []string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		favFlavors: []string{
			"Martini",
			"Shaken",
			"Not stirred",
		},
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		favFlavors: []string{
			"James",
			"Raspberry",
		},
	}
	people := []person{p1, p2}
	jsonBlob, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBlob))
}
