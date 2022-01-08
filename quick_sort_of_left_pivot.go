/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

// This Quick sort is implemented with left element selected as the pivot

package main

func QuickSortLP(a []int) {
	quickSortLP(a, 0, len(a)-1)
}

func quickSortLP(a []int, lo, hi int) {

	if lo >= hi {
		return
	}
	pivot := partitionLP(a, lo, hi)
	quickSortLP(a, lo, pivot-1)
	quickSortLP(a, pivot+1, hi)
}

func partitionLP(a []int, left, right int) int {
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
