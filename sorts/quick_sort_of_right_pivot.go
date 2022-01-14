/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

// This Quick sort is implemented with right element selected as the pivot

package sorts

func QuickSortRP(a []int) {
	quickSortLP(a, 0, len(a))
}

func quickSortRP(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}
	pivot := partitionRP(a, lo, hi-1)
	quickSortRP(a, lo, pivot)
	quickSortRP(a, pivot+1, hi)
}

func partitionRP(a []int, left, right int) int {
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
