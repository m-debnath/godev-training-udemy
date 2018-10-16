package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	jsonStr := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	fmt.Println(jsonStr)
	var people []person
	err := json.Unmarshal([]byte(jsonStr), &people)
	if err != nil {
		fmt.Println("Error converting JSON:", `json.Unmarshal([]byte(jsonStr), &people)`)
		fmt.Println("jsonStr:", jsonStr)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Println("Person number", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
	}
}
