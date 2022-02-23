/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package simplytest

import (
	"fmt"
	"strings"
	"time"

	"github.com/logrusorgru/aurora/v3"
)

type OutputForm struct {
	active bool
	name   string
	Time   int64
	equl   bool
	err    string
}

type Queue []OutputForm

func callSortTest(origin, verify []int) {
	var q Queue
	q = append(q, CallBubbleSort(origin, verify, "bubble sort"))
	q = append(q, CallCocktailSort(origin, verify, "cocktail sort"))
	q = append(q, CallInsertionSort(origin, verify, "insertion sort"))
	q = append(q, CallSelectionSort(origin, verify, "selection sort"))
	q = append(q, CallShellSort(origin, verify, "shell sort"))
	q = append(q, CallBottomUpMergeSort(origin, verify, "bottom-up merge sort"))
	q = append(q, CallTopDownMergeSort(origin, verify, "top-down merge sort"))
	q = append(q, CallParallelMergeSort(origin, verify, "parallel merge sort"))
	q = append(q, CallHeapSort(origin, verify, "heap sort"))
	q = append(q, CallLPQuickSort(origin, verify, "left-pivot quick sort"))
	q = append(q, CallQuickSort(origin, verify, "middle-pivot quick sort"))
	q = append(q, CallRPQuickSort(origin, verify, "right-pivot quick sort"))
	q = append(q, CallParallelLPQuickSort(origin, verify, "parallel left-pivot quick sort"))
	q = append(q, CallParallelMPQuickSort(origin, verify, "parallel middle-pivot quick sort"))
	q = append(q, CallParallelRPQuickSort(origin, verify, "parallel right-pivot quick sort"))
	q = append(q, CallDualPivotQuickSort(origin, verify, "dual-pivot quick sort"))
	q = append(q, CallParallelDualPivotQuickSort(origin, verify, "parallel dual-pivot quick sort"))
	q = append(q, CallBinaryInsertionSort(origin, verify, "binary insertion sort"))
	q = append(q, CallTimSort(origin, verify, "tim sort"))
	q = append(q, CallBitonicSort(origin, verify, "bitonic sort"))
	q = append(q, CallParallelBitonicSort(origin, verify, "parallel bitonic sort"))
	q = append(q, CallIntroSort(origin, verify, "intro sort"))
	q = append(q, CallParallelIntroSort(origin, verify, "parallel intro sort"))
	q = append(q, CallCycleSort(origin, verify, "cycle sort"))
	q = append(q, CallOddEvenSort(origin, verify, "odd-even sort"))
	q = append(q, CallOddEvenMergeSort(origin, verify, "odd-even merge sort"))
	var pf string = ""

	pf += fmt.Sprintf("\n+%s+\n", strings.Repeat("-", 97))
	pf += fmt.Sprintf("| %35s | %23s | %18s | %10s | \t %s \n", "name", "ns", "ms", "verify", aurora.BrightRed("(err mag)"))

	for _, v := range q {
		if v.active {
			pf += fmt.Sprintf("|%s|\n", strings.Repeat("-", 97))
			pf += fmt.Sprintf("| %35s | %20d ns | %15d ms | %10t |%s\n", v.name, v.Time, v.Time/int64(time.Millisecond), v.equl, v.err)
		}
	}

	pf += fmt.Sprintf("+%s+\n\n\n", strings.Repeat("-", 97))

	fmt.Print(pf)
}
