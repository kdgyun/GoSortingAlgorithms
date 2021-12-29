// This Parallel Quick sort is implemented with left element selected as the pivot

package main

import (
	"runtime"
	"sync"
)

func ParallelQuickSortLP(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelQuickSortLP(a, 0, len(a)-1)
}

func parallelQuickSortLP(a []int, lo, hi int) {

	if lo >= hi {
		return
	}

	// 2048 is threshold
	if hi-lo > 2048 {
		pivot := parallelPartitionLP(a, lo, hi)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			parallelQuickSortLP(a, lo, pivot-1)
		}()
		parallelQuickSortLP(a, pivot+1, hi)
		wg.Wait()
	} else {
		defaultQuickSortLP(a, lo, hi)
	}
}

func parallelPartitionLP(a []int, left, right int) int {
	lo := left
	hi := right
	pivot := a[left]

	for lo < hi {
		for a[hi] > pivot {
			hi--
		}
		for a[lo] <= pivot && lo < hi {
			lo++
		}

		a[lo], a[hi] = a[hi], a[lo]
	}
	a[left], a[lo] = a[lo], a[left]
	return lo
}

func defaultQuickSortLP(a []int, lo, hi int) {

	if lo >= hi {
		return
	}
	pivot := parallelPartitionLP(a, lo, hi)
	defaultQuickSortLP(a, lo, pivot-1)
	defaultQuickSortLP(a, pivot+1, hi)
}
