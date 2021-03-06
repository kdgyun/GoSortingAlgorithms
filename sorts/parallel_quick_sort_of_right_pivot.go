/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

// This Quick sort is implemented with right element selected as the pivot

package sorts

import (
	"runtime"
	"sync"
)

func ParallelQuickSortRP(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelQuickSortRP(a, 0, len(a))
}

func parallelQuickSortRP(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := parallelPartitionRP(a, lo, hi-1)
	// 2048 is threshold
	if hi-lo > 2048 {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			parallelQuickSortRP(a, lo, pivot)
		}()
		parallelQuickSortRP(a, pivot+1, hi)
		wg.Wait()

	} else {
		basicQuickSortRP(a, lo, pivot)
		basicQuickSortRP(a, pivot+1, hi)
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

func basicQuickSortRP(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := parallelPartitionRP(a, lo, hi-1)
	basicQuickSortRP(a, lo, pivot)
	basicQuickSortRP(a, pivot+1, hi)
}
