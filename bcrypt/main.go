package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(string(bs))
	s1 := `password1234`
	bs1, err1 := bcrypt.GenerateFromPassword([]byte(s1), bcrypt.MinCost)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(s1)
	fmt.Println(string(bs1))

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs1, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
