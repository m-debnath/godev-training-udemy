package main

import "testing"

func TestSubtract(t *testing.T) {
	v := subtract(42.0, 2.0)
	if v != 40.0 {
		t.Error("Expected 40.0 Got", v)
	}
}
