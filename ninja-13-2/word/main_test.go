package word

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/m-debnath/godev-training-udemy/ninja-13-2/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("I am Mukul. I am also Debnath.")
	ans := map[string]int{
		"I":        2,
		"am":       2,
		"Mukul.":   1,
		"also":     1,
		"Debnath.": 1,
	}
	eq := reflect.DeepEqual(m, ans)
	if !eq {
		t.Error("Expected", ans, "Got", m)
	}
}

func TestCount(t *testing.T) {
	i := Count("Mukul")
	if i != 5 {
		t.Error("Expected", 5, "Got", i)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("Mukul"))
	// Output:
	// 5
}

func ExampleUseCount() {
	fmt.Println(UseCount("I am Mukul. I am also Debnath."))
	// Output:
	// map[I:2 am:2 Mukul.:1 also:1 Debnath.:1]
}
