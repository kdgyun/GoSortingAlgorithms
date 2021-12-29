package main

import (
	"fmt"
	"time"
)

func CallBubbleSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	BubbleSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallInsertionSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	InsertionSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallSelectionSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	SelectionSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallShellSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	ShellSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallBottomUpMergeSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	BottomUpMergeSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallTopDownMergeSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	TopDownMergeSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallParallelMergeSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	ParallelMergeSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallHeapSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	HeapSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallLPQuickSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	QuickSortLP(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallQuickSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	QuickSort(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallRPQuickSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	QuickSortRP(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}

func CallParallelLPQuickSort(origin []int, verify []int, callName string) (int64, bool, string) {
	test := make([]int, len(origin))
	copy(test, origin)
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	ParallelQuickSortLP(test)
	end := time.Since(start)
	eq, err := Equal(verify, test)
	return end.Nanoseconds(), eq, err
}
