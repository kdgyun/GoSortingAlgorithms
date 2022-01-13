/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

import (
	"fmt"
	"time"

	"GoSortingAlgorithms/sorts"
)

func CallBubbleSort(origin []int, verify []int, callName string) OutputForm {
	if BUBBLE_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.BubbleSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallCocktailSort(origin []int, verify []int, callName string) OutputForm {
	if COCKTAIL_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.CocktailSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallInsertionSort(origin []int, verify []int, callName string) OutputForm {
	if INSERTION_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.InsertionSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallSelectionSort(origin []int, verify []int, callName string) OutputForm {
	if SELECTION_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.SelectionSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallShellSort(origin []int, verify []int, callName string) OutputForm {
	if SHELL_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ShellSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallBottomUpMergeSort(origin []int, verify []int, callName string) OutputForm {
	if BOTTOM_UP_MERGE_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.BottomUpMergeSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallTopDownMergeSort(origin []int, verify []int, callName string) OutputForm {
	if TOP_DOWN_MERGE_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.TopDownMergeSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallParallelMergeSort(origin []int, verify []int, callName string) OutputForm {
	if PARALLEL_MERGE_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ParallelMergeSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallHeapSort(origin []int, verify []int, callName string) OutputForm {
	if HEAP_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.HeapSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallLPQuickSort(origin []int, verify []int, callName string) OutputForm {
	if LEFT_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.QuickSortLP(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallQuickSort(origin []int, verify []int, callName string) OutputForm {
	if MIDDLE_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.QuickSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallRPQuickSort(origin []int, verify []int, callName string) OutputForm {
	if RIGHT_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.QuickSortRP(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallParallelLPQuickSort(origin []int, verify []int, callName string) OutputForm {
	if PARALLEL_LEFT_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ParallelQuickSortLP(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallParallelMPQuickSort(origin []int, verify []int, callName string) OutputForm {
	if PARALLEL_MIDDLE_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ParallelQuickSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallParallelRPQuickSort(origin []int, verify []int, callName string) OutputForm {
	if PARALLEL_RIGHT_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ParallelQuickSortRP(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallDualPivotQuickSort(origin []int, verify []int, callName string) OutputForm {
	if DUAL_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.DualPivotQuickSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallParallelDualPivotQuickSort(origin []int, verify []int, callName string) OutputForm {
	if PARALLEL_DUAL_PIVOT_QUICK_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ParallelDualPivotQuickSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallBinaryInsertionSort(origin []int, verify []int, callName string) OutputForm {
	if BINARY_INSERTION_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.BinarySort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallTimSort(origin []int, verify []int, callName string) OutputForm {
	if TIM_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.TimSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallBitonicSort(origin []int, verify []int, callName string) OutputForm {
	if BITONIC_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.BitonicSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}

func CallParallelBitonicSort(origin []int, verify []int, callName string) OutputForm {
	if PARALLEL_BITONIC_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		sorts.ParallelBitonicSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}
