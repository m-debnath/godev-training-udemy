package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{42, 9, 1}, 52},
		test{[]int{-1, 6}, 5},
		test{[]int{7, 4, 2, 9}, 22},
	}
	for _, v := range tests {
		ans := mySum(v.data...)
		if ans != v.answer {
			t.Error("Expected", v.answer, "Got", ans)
		}
	}
}
