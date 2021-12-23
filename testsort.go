package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
)

func test() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	var verify []int
	iter := rand.Intn(100000)
	for i := 0; i < iter; i++ {
		verify = append(verify, rand.Int())
	}
	bubble := make([]int, len(verify))
	insertion := make([]int, len(verify))
	selection := make([]int, len(verify))
	shell := make([]int, len(verify))
	merge := make([]int, len(verify))
	heap := make([]int, len(verify))

	copy(bubble, verify)
	copy(insertion, verify)
	copy(selection, verify)
	copy(shell, verify)
	copy(merge, verify)
	copy(heap, verify)

	sort.Ints(verify)

	BubbleSort(bubble)
	InsertionSort(insertion)
	SelectionSort(selection)
	ShellSort(shell)
	MergeSort(merge)
	HeapSort(heap)

	fmt.Printf("slice length : %d\n", len(verify))
	fmt.Printf("bubble sort : %t\n", Equal(verify, bubble))
	fmt.Printf("insertion sort : %t\n", Equal(verify, insertion))
	fmt.Printf("selection sort : %t\n", Equal(verify, selection))
	fmt.Printf("shell sort : %t\n", Equal(verify, shell))
	fmt.Printf("merge sort : %t\n", Equal(verify, merge))
	fmt.Printf("heap sort : %t\n", Equal(verify, heap))
}

func main() {

	for i := 1; i < 11; i++ {
		fmt.Printf("\n===========[test%d]===========\n", i)
		test()
	}

}

func Equal(verify, b []int) bool {
	if len(verify) != len(b) {
		return false
	}
	for i, v := range verify {
		if v != b[i] {
			return false
		}
	}
	return true
}
