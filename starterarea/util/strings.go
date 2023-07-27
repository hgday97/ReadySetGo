package util

var greeting = "Hello!" // Package scope, so can be accessed by anything in util

func StringLength(s string) int {
	return len(s)
}
