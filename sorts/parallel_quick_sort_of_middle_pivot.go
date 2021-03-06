/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

// This Quick sort is implemented with middle element selected as the pivot

package sorts

import (
	"runtime"
	"sync"
)

func ParallelQuickSort(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelQuickSort(a, 0, len(a))
}

func parallelQuickSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := parallelPartition(a, lo, hi-1)
	// 2048 is threshold
	if hi-lo > 2048 {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			parallelQuickSort(a, lo, pivot)

		}()
		parallelQuickSort(a, pivot+1, hi)
		wg.Wait()
	} else {
		basicQuickSort(a, lo, pivot)
		basicQuickSort(a, pivot+1, hi)
	}
}

func parallelPartition(a []int, left, right int) int {
	lo := left
	hi := right
	pivot := a[(left+right)>>1]

	for true {
		for a[lo] < pivot {
			lo++
		}
		for a[hi] > pivot && lo <= hi {
			hi--
		}
		if lo >= hi {
			return hi
		}

		a[lo], a[hi] = a[hi], a[lo]
	}
	return hi
}

func basicQuickSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := parallelPartition(a, lo, hi-1)
	basicQuickSort(a, lo, pivot)
	basicQuickSort(a, pivot+1, hi)
}
