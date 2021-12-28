// This Quick sort is implemented with middle element selected as the pivot

package main

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {

	if lo >= hi {
		return
	}
	pivot := partition(a, lo, hi)
	quickSort(a, lo, pivot)
	quickSort(a, pivot+1, hi)
}

func partition(a []int, left, right int) int {
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
