package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func test() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	var verify []int
	iter := rand.Intn(300000)
	fmt.Printf("[array length : %d]\n", iter)

	fmt.Println("make arrays...")
	for i := 0; i < iter; i++ {
		verify = append(verify, rand.Int())
	}
	bubble := make([]int, len(verify))
	insertion := make([]int, len(verify))
	selection := make([]int, len(verify))
	shell := make([]int, len(verify))
	bmerge := make([]int, len(verify))
	tmerge := make([]int, len(verify))
	heap := make([]int, len(verify))

	fmt.Println("copy arrays...")
	copy(bubble, verify)
	copy(insertion, verify)
	copy(selection, verify)
	copy(shell, verify)
	copy(bmerge, verify)
	copy(tmerge, verify)
	copy(heap, verify)

	fmt.Println("runing default sort...")
	sort.Ints(verify)

	t1 := CallBubbleSort(bubble, "bubble sort")
	t2 := CallInsertionSort(insertion, "insertion sort")
	t3 := CallSelectionSort(selection, "selection sort")
	t4 := CallShellSort(shell, "shell sort")
	t5 := CallBottomUpMergeSort(bmerge, "bottom-up merge sort")
	t6 := CallTopDownMergeSort(tmerge, "top-down merge sort")
	t7 := CallHeapSort(heap, "heap sort")

	var pf string = ""
	var s1 bool
	var ef string

	pf += fmt.Sprintf("\n+%s+\n", strings.Repeat("-", 82))
	pf += fmt.Sprintf("| %25s | %18s | %18s | %10s | \t (err mag) \n", "name", "ns", "ms", "verify")
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, bubble)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "bubble sort", t1, t1/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, insertion)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "insertion sort", t2, t2/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, selection)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "selection sort", t3, t3/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, shell)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "shell sort", t4, t4/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, bmerge)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "bottom-up merge sort", t5, t5/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, tmerge)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "top-down merge sort", t6, t6/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	s1, ef = Equal(verify, heap)
	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "heap sort", t7, t7/int64(time.Millisecond), s1, ef)
	pf += fmt.Sprintf("+%s+\n\n\n", strings.Repeat("-", 82))

	fmt.Print(pf)

}
func main() {

	for i := 1; i < 11; i++ {
		fmt.Printf("\n===========[test%d]===========\n", i)
		test()

	}

}
