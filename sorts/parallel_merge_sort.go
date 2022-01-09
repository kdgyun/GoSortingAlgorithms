/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

import (
	"runtime"
	"sync"
)

func ParallelMergeSort(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	temp := make([]int, len(a))
	parallelMergeSort(a, temp, 0, len(a))
}

func parallelMergeSort(a, temp []int, left int, right int) {

	if right-left < 2 {
		return
	}

	// 2048 is threshold
	if right-left > 2048 { // parallel

		mid := (left + right) >> 1

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			parallelMergeSort(a, temp, left, mid)
		}()

		parallelMergeSort(a, temp, mid, right)
		wg.Wait()
		mergeForPal(a, temp, left, mid, right)
	} else {
		basicMergeSort(a, temp, left, right)
	}
}

func basicMergeSort(a, temp []int, left int, right int) {

	if right-left < 2 {
		return
	}
	mid := (left + right) / 2
	topDownMergeSort(a, temp, left, mid)
	topDownMergeSort(a, temp, mid, right)
	mergeForPal(a, temp, left, mid, right)
}

func mergeForPal(origin, temp []int, left, mid, right int) {
	l := left
	r := mid
	idx := left
	for l < mid && r < right {
		if origin[l] <= origin[r] {
			temp[idx] = origin[l]
			idx++
			l++
		} else {
			temp[idx] = origin[r]
			idx++
			r++
		}
	}
	if l == mid {
		for r < right {
			temp[idx] = origin[r]
			idx++
			r++
		}
	} else {
		for l < mid {
			temp[idx] = origin[l]
			idx++
			l++
		}
	}

	copy(origin[left:], temp[left:right])
}
