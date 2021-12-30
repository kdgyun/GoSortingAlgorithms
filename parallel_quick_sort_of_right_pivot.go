// This Quick sort is implemented with right element selected as the pivot

package main

import (
	"runtime"
	"sync"
)

func ParallelQuickSortRP(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	quickSortLP(a, 0, len(a)-1)
}

func parallelQuickSortRP(a []int, lo, hi int) {

	if lo >= hi {
		return
	}

	// 2048 is threshold
	if hi-lo > 2048 {
		var wg sync.WaitGroup
		wg.Add(1)
		pivot := parallelPartitionRP(a, lo, hi)
		go func() {
			defer wg.Done()
			parallelQuickSortRP(a, lo, pivot-1)
		}()
		parallelQuickSortRP(a, pivot+1, hi)
		wg.Wait()

	} else {
		defaultQuickSortRP(a, lo, hi)
	}

}

func parallelPartitionRP(a []int, left, right int) int {
	lo := left
	hi := right
	pivot := a[right]

	for lo < hi {
		for a[lo] < pivot && lo < hi {
			lo++
		}
		for a[hi] >= pivot && lo < hi {
			hi--
		}

		a[lo], a[hi] = a[hi], a[lo]
	}
	a[right], a[hi] = a[hi], a[right]
	return lo
}

func defaultQuickSortRP(a []int, lo, hi int) {

	if lo >= hi {
		return
	}
	pivot := parallelPartitionRP(a, lo, hi)
	defaultQuickSortRP(a, lo, pivot-1)
	defaultQuickSortRP(a, pivot+1, hi)
}
