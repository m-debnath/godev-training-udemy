// Package acdc asks you if you're ready to rock.
package acdc

// Sum returns the sum of unlimited values of type int.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
