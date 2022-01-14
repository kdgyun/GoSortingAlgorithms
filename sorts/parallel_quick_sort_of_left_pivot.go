/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

// This Parallel Quick sort is implemented with left element selected as the pivot

package sorts

import (
	"runtime"
	"sync"
)

func ParallelQuickSortLP(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelQuickSortLP(a, 0, len(a))
}

func parallelQuickSortLP(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}

	// 2048 is threshold
	if hi-lo > 2048 {
		pivot := parallelPartitionLP(a, lo, hi-1)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			parallelQuickSortLP(a, lo, pivot)
		}()
		parallelQuickSortLP(a, pivot+1, hi)
		wg.Wait()
	} else {
		basicQuickSortLP(a, lo, hi)
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

func basicQuickSortLP(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := parallelPartitionLP(a, lo, hi-1)
	basicQuickSortLP(a, lo, pivot)
	basicQuickSortLP(a, pivot+1, hi)
}
