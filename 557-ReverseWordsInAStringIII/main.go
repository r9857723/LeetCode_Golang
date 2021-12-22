package main

import (
	"fmt"
	"strings"
)

func main() {

	q := "Let's take LeetCode contest"
	fmt.Println(reverseWords(q))
}

func reverseWords(s string) string {
	rtn := ""
	for _, g := range strings.Split(s," ") {
		str := ""
		for i := len(g) - 1 ; i >= 0 ; i-- {
			str += string(g[i])
		}
			rtn += " " + str
	}
	return strings.TrimSpace(rtn)
}