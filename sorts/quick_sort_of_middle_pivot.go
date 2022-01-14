/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

// This Quick sort is implemented with middle element selected as the pivot

package sorts

func QuickSort(a []int) {
	quickSort(a, 0, len(a))
}

func quickSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := partition(a, lo, hi-1)
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
