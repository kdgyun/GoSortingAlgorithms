package main

import (
	"fmt"
	"strings"
)

func RandomTest() {
	fmt.Printf("\n+%s+\n", strings.Repeat("=", 82))
	fmt.Printf("|%46s%36s|\n", "Random Test", " ")
	fmt.Printf("+%s+\n\n\n", strings.Repeat("=", 82))
	for _, v := range lengths {
		runRandom(v)
	}
}
func runRandom(len int) {
	origin, verify := RANDOM_INT(len)
	run(origin, verify)
}
