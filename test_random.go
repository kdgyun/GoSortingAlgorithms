package main

import (
	"fmt"
	"strings"
	"time"
)

func RandomTest() {
	fmt.Printf("\n+%s+\n", strings.Repeat("=", 82))
	fmt.Printf("|%46s%36s|\n", "Random Test", " ")
	fmt.Printf("+%s+\n\n\n", strings.Repeat("=", 82))
	for _, v := range length {
		runRandom(v)
	}
}
func runRandom(len int) {

	origin, verify := RANDOM_INT(len)

	t1, eq1, err1 := CallBubbleSort(origin, verify, "bubble sort")
	t2, eq2, err2 := CallInsertionSort(origin, verify, "insertion sort")
	t3, eq3, err3 := CallSelectionSort(origin, verify, "selection sort")
	t4, eq4, err4 := CallShellSort(origin, verify, "shell sort")
	t5, eq5, err5 := CallBottomUpMergeSort(origin, verify, "bottom-up merge sort")
	t6, eq6, err6 := CallTopDownMergeSort(origin, verify, "top-down merge sort")
	t7, eq7, err7 := CallParallelMergeSort(origin, verify, "parallel merge sort")
	t8, eq8, err8 := CallHeapSort(origin, verify, "heap sort")

	var pf string = ""

	pf += fmt.Sprintf("\n+%s+\n", strings.Repeat("-", 82))
	pf += fmt.Sprintf("| %25s | %18s | %18s | %10s | \t (err mag) \n", "name", "ns", "ms", "verify")
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "bubble sort", t1, t1/int64(time.Millisecond), eq1, err1)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "insertion sort", t2, t2/int64(time.Millisecond), eq2, err2)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "selection sort", t3, t3/int64(time.Millisecond), eq3, err3)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "shell sort", t4, t4/int64(time.Millisecond), eq4, err4)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "bottom-up merge sort", t5, t5/int64(time.Millisecond), eq5, err5)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "top-down merge sort", t6, t6/int64(time.Millisecond), eq6, err6)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "parallel merge sort", t7, t7/int64(time.Millisecond), eq7, err7)
	pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 82))

	pf += fmt.Sprintf("| %25s | %15d ns | %15d ms | %10t |%s\n", "heap sort", t8, t8/int64(time.Millisecond), eq8, err8)
	pf += fmt.Sprintf("+%s+\n\n\n", strings.Repeat("-", 82))

	fmt.Print(pf)
}
