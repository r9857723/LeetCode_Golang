package main

import (
	"strings"
)
func main() {
	str := "Hello, my name is John"
	countSegments(str)
}
func countSegments(s string) int {
	return len(strings.Fields(s))
}