package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	i := Years(3)
	if i != 21 {
		t.Error("Expected", 21, "Got", i)
	}
}
func TestYearsTwo(t *testing.T) {
	i := YearsTwo(3)
	if i != 21 {
		t.Error("Expected", 21, "Got", i)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(3)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	// Output:
	// 28
}
