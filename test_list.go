package main

import (
	"fmt"
	"time"
)

func CallBubbleSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	BubbleSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallInsertionSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	InsertionSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallSelectionSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	SelectionSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallShellSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	ShellSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallBottomUpMergeSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	BottomUpMergeSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallTopDownMergeSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	TopDownMergeSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallParallelMergeSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	ParallelMergeSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}

func CallHeapSort(a []int, callName string) int64 {
	fmt.Printf("runing %s...\n", callName)
	start := time.Now()
	HeapSort(a)
	end := time.Since(start)
	return end.Nanoseconds()
}
