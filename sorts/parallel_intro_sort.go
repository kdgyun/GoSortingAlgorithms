package sorts

import (
	"math"
	"runtime"
	"sync"
)

const (
	SIZE_THREADHOLD   = 16
	T_SIZE_THREADHOLD = 2048
)

func ParallelIntroSort(a []int) {
	parallelIntroSort(a, 0, len(a))
}

func parallelIntroSort(a []int, lo, hi int) {
	if hi-lo <= SIZE_THREADHOLD {
		__basicInsertionSort(a, lo, hi)
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	limit := 2 * int(math.Log2(float64(hi-lo)))
	__basicParallelQuickSort(a, lo, hi, limit)
	__basicInsertionSort(a, lo, hi)
}

func __basicInsertionSort(a []int, lo, hi int) {

	for i := lo + 1; i < hi; i++ {
		target := a[i]
		j := i - 1

		for j >= lo && target < a[j] {
			a[j+1] = a[j]
			j--
		}

		a[j+1] = target
	}
}

func __basicHeapSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	last := hi - 1
	parent := ((last - 1) >> 1)

	for parent >= lo {
		__basicHeapify(a, parent, last)
		parent--
	}

	end := hi - 1
	for end > lo {
		a[lo], a[end] = a[end], a[lo]
		end--
		__basicHeapify(a, lo, end)
	}
}

func __basicHeapify(a []int, cur, last int) {
	var left, right, large int

	for ((cur << 1) + 1) <= last {
		left = (cur << 1) + 1
		right = (cur << 1) + 2
		large = cur

		if a[left] > a[large] {
			large = left
		}

		if right <= last && a[right] > a[large] {
			large = right
		}

		if large != cur {
			a[cur], a[large] = a[large], a[cur]
			cur = large
		} else {
			return
		}
	}
}

func __basicParallelQuickSort(a []int, lo, hi, depthLimit int) {

	if hi-lo <= SIZE_THREADHOLD {
		__basicInsertionSort(a, lo, hi)
		return
	}
	if depthLimit == 0 {
		__basicHeapSort(a, lo, hi)
		return
	}
	pivot := __basicPartition(a, lo, hi-1)
	if hi-lo > T_SIZE_THREADHOLD {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			__basicParallelQuickSort(a, lo, pivot, depthLimit-1)
		}()
		__basicParallelQuickSort(a, pivot, hi, depthLimit-1)
		wg.Wait()
	} else {
		__basicQuickSort(a, lo, pivot, depthLimit-1)
		__basicQuickSort(a, pivot, hi, depthLimit-1)
	}
}

func __basicQuickSort(a []int, lo, hi, depthLimit int) {

	if hi-lo <= SIZE_THREADHOLD {
		__basicInsertionSort(a, lo, hi)
		return
	}
	if depthLimit == 0 {
		__basicHeapSort(a, lo, hi)
		return
	}
	pivot := __basicPartition(a, lo, hi-1)
	__basicQuickSort(a, lo, pivot, depthLimit-1)
	__basicQuickSort(a, pivot, hi, depthLimit-1)
}

func __basicPartition(a []int, left, right int) int {
	lo := left
	hi := right
	pivot := a[left+((right-left)>>1)]

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
