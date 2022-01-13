package sorts

import (
	"runtime"
	"sync"
)

func ParallelDualPivotQuickSort(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelDualPivotQuickSort(a, 0, len(a))

}

func parallelDualPivotQuickSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}

	// 4096 is threshold
	if hi-lo > 4096 {
		var wg sync.WaitGroup
		lp, rp := parallel3wayQuickPartition(a, lo, hi-1)
		wg.Add(2)
		go func() {
			defer wg.Done()
			parallelDualPivotQuickSort(a, lo, lp)
		}()
		go func() {
			defer wg.Done()
			parallelDualPivotQuickSort(a, lp+1, rp)
		}()
		parallelDualPivotQuickSort(a, rp+1, hi)
		wg.Wait()
	} else {
		basicDualPivotQuickSort(a, lo, hi)
	}

}

func basicDualPivotQuickSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}

	lp, rp := parallel3wayQuickPartition(a, lo, hi-1)
	basicDualPivotQuickSort(a, lo, lp)
	basicDualPivotQuickSort(a, lp+1, rp)
	basicDualPivotQuickSort(a, rp+1, hi)

}

func parallel3wayQuickPartition(a []int, lo, hi int) (int, int) {

	if a[lo] > a[hi] {
		a[lo], a[hi] = a[hi], a[lo]
	}

	l := lo + 1
	k := lo + 1
	g := hi - 1

	leftPivot := a[lo]
	rightPivot := a[hi]

	for k <= g {
		if a[k] < leftPivot {
			a[k], a[l] = a[l], a[k]
			l++
		} else if a[k] >= rightPivot {
			for rightPivot < a[g] && k < g {
				g--
			}
			a[k], a[g] = a[g], a[k]
			g--

			if a[k] < leftPivot {
				a[k], a[l] = a[l], a[k]
				l++
			}
		}
		k++
	}
	l--
	g++

	a[lo], a[l] = a[l], a[lo]
	a[hi], a[g] = a[g], a[hi]

	return l, g
}
