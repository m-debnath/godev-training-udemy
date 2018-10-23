package greeting

// Greet returns a greeting when passed an unlimited values of type string.
func Greet(s ...string) string {
	greeting := "Welcome to our program"
	for _, v := range s {
		greeting += " " + v
	}
	return greeting
}
