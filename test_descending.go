package main

import (
	"fmt"
	"strings"
)

func DescendingTest() {
	fmt.Printf("\n+%s+\n", strings.Repeat("=", 82))
	fmt.Printf("|%49s%33s|\n", "Descending Test", " ")
	fmt.Printf("+%s+\n\n\n", strings.Repeat("=", 82))
	for _, v := range lengths {
		runDescending(v)
	}
}
func runDescending(len int) {

	origin, verify := DESCENDING_INT(len)

	run(origin, verify)
}
