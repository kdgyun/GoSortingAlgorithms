package sorts

import (
	"math"
)

const (
	S_THREADHOLD = 16
)

func IntroSort(a []int) {
	introSort(a, 0, len(a))
}

func introSort(a []int, lo, hi int) {
	if hi-lo <= S_THREADHOLD {
		__insertionSort(a, lo, hi)
		return
	}
	limit := 2 * int(math.Log2(float64(hi-lo)))
	__quickSort(a, lo, hi, limit)
	__insertionSort(a, lo, hi)
}

func __insertionSort(a []int, lo, hi int) {

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

func __heapSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	last := hi - 1
	parent := ((last - 1) >> 1)

	for parent >= lo {
		__heapify(a, parent, last)
		parent--
	}

	end := hi - 1
	for end > lo {
		a[lo], a[end] = a[end], a[lo]
		end--
		__heapify(a, lo, end)
	}
}

func __heapify(a []int, cur, last int) {
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

func __quickSort(a []int, lo, hi, depthLimit int) {

	if hi-lo <= S_THREADHOLD {
		__insertionSort(a, lo, hi)
		return
	}
	if depthLimit == 0 {
		__heapSort(a, lo, hi)
		return
	}
	pivot := __partition(a, lo, hi-1)
	__quickSort(a, lo, pivot, depthLimit-1)
	__quickSort(a, pivot, hi, depthLimit-1)
}

func __partition(a []int, left, right int) int {
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
