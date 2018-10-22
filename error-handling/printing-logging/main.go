package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	// f, err := os.Create("log.txt")
	// if err != nil {
	// 	fmt.Println(err, "happened while creating log file")
	// }
	// defer f.Close()
	// log.SetOutput(f)

	// f2, err := os.Open("no-file.txt")
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("error:", err)
		// log.Println("error:", err)
		// log.Fatalln("error:", err)
		log.Panicln("error:", err)
	}
	// defer f2.Close()
}

func foo() {
	fmt.Println("When os.Exit is called deferred functions don't run")
}
