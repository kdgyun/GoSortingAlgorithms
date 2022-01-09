/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

import (
	"fmt"
	"time"
)

func CallBubbleSort(origin []int, verify []int, callName string) OutputForm {
	if BUBBLE_SORT {
		test := make([]int, len(origin))
		copy(test, origin)
		fmt.Printf("runing %s...\n", callName)
		start := time.Now()
		BubbleSort(test)
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
		CocktailSort(test)
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
		InsertionSort(test)
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
		SelectionSort(test)
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
		ShellSort(test)
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
		BottomUpMergeSort(test)
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
		TopDownMergeSort(test)
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
		ParallelMergeSort(test)
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
		HeapSort(test)
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
		QuickSortLP(test)
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
		QuickSort(test)
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
		QuickSortRP(test)
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
		ParallelQuickSortLP(test)
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
		ParallelQuickSort(test)
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
		ParallelQuickSortRP(test)
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
		BinarySort(test)
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
		Timsort(test)
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
		BitonicSort(test)
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
		ParallelBitonicSort(test)
		end := time.Since(start)
		eq, err := Equal(verify, test)
		return OutputForm{true, callName, end.Nanoseconds(), eq, err}
	}
	return OutputForm{false, callName, -1, false, ""}
}
