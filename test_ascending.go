package main

import (
	"fmt"
	"strings"
)

func AscendingTest() {
	fmt.Printf("\n+%s+\n", strings.Repeat("=", 82))
	fmt.Printf("|%48s%34s|\n", "Ascending Test", " ")
	fmt.Printf("+%s+\n\n\n", strings.Repeat("=", 82))
	for _, v := range length {
		runAscending(v)
	}
}

func runAscending(len int) {
	origin, verify := ASCENDING_INT(len)

	run(origin, verify)
}
