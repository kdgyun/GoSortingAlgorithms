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

func ParallelOddEvenMergeSort(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelOddEvenMergeSort(a, 0, len(a))
}

func parallelOddEvenMergeSort(a []int, lo, hi int) {
	if hi-lo > 1 {
		mid := (hi + lo) >> 1
		// 2048 is threshold
		if hi-lo > 2048 {

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				parallelOddEvenMergeSort(a, lo, mid)
			}()
			parallelOddEvenMergeSort(a, mid, hi)
			wg.Wait()
			parallelOddEvenMerge(a, lo, hi, 1)
		} else {
			basicOddEvenMergeSort(a, lo, mid)
			basicOddEvenMergeSort(a, mid, hi)
			basicOddEvenMerge(a, lo, hi, 1)
		}
	}
}

func parallelOddEvenMerge(a []int, lo, hi, dist int) {
	subDist := dist << 1
	if subDist < (hi - lo) {

		if subDist < 32 {
			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				parallelOddEvenMerge(a, lo, hi, subDist)
			}()
			parallelOddEvenMerge(a, lo+dist, hi, subDist)
			wg.Wait()
		} else {
			basicOddEvenMerge(a, lo, hi, subDist)
			basicOddEvenMerge(a, lo+dist, hi, subDist)
		}
		for i := lo + dist; i+dist < hi; i += subDist {

			if a[i] > a[i+dist] {
				a[i], a[i+dist] = a[i+dist], a[i]
			}
		}
	} else {
		if a[lo] > a[lo+dist] {
			a[lo], a[lo+dist] = a[lo+dist], a[lo]
		}
	}
}

func basicOddEvenMergeSort(a []int, lo, hi int) {
	if hi-lo > 1 {
		mid := (hi + lo) >> 1
		basicOddEvenMergeSort(a, lo, mid)
		basicOddEvenMergeSort(a, mid, hi)
		basicOddEvenMerge(a, lo, hi, 1)
	}
}

func basicOddEvenMerge(a []int, lo, hi, dist int) {
	subDist := dist << 1
	if subDist < (hi - lo) {
		basicOddEvenMerge(a, lo, hi, subDist)
		basicOddEvenMerge(a, lo+dist, hi, subDist)
		for i := lo + dist; i+dist < hi; i += subDist {

			if a[i] > a[i+dist] {
				a[i], a[i+dist] = a[i+dist], a[i]
			}
		}
	} else {
		if a[lo] > a[lo+dist] {
			a[lo], a[lo+dist] = a[lo+dist], a[lo]
		}
	}
}
