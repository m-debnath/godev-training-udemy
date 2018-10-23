package greeting

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Mukul", "Debnath")
	if s != "Welcome to our program Mukul Debnath" {
		t.Error("Expected", "Welcome to our program Mukul Debnath", "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Mukul", "Debnath"))
	// Output:
	// Welcome to our program Mukul Debnath
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Mukul", "Debnath")
	}
}
