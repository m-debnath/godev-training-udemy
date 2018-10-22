package main

import (
	"fmt"
	"log"
)

type customErr struct {
	err string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Custom error: %v", c.err)
}

func main() {
	c := customErr{
		err: "Custom Error 1",
	}

	foo(c)
}

func foo(e error) {
	//fmt.Println(c.Error())
	log.Panicln(e)
}
