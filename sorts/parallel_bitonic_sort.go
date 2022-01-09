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

func parallelBitonicMerge(a []int, lo, len int, dir bool) {

	if len > 1 {
		m := basicGreatestPowerOfTwoLessThan(len)
		for i := lo; i < lo+len-m; i++ {
			if dir == (a[i] > a[i+m]) {
				a[i], a[i+m] = a[i+m], a[i]
			}
		}
		// 2048 is threshold
		if len > 2048 { // parallel

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				parallelBitonicMerge(a, lo, m, dir)
			}()
			parallelBitonicMerge(a, lo+m, len-m, dir)
			wg.Wait()
		} else {
			basicBitonicMerge(a, lo, m, dir)
			basicBitonicMerge(a, lo+m, len-m, dir)
		}

	}

}

func ParallelBitonicSort(a []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parallelBitonicSort(a, 0, len(a), true)
}

func parallelBitonicSort(a []int, lo, len int, dir bool) {

	if len > 1 {
		mid := len >> 1

		// 2048 is threshold
		if len > 2048 { // parallel

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				parallelBitonicSort(a, lo, mid, !dir)
			}()

			parallelBitonicSort(a, lo+mid, len-mid, dir)
			wg.Wait()
			parallelBitonicMerge(a, lo, len, dir)
		} else {
			basicBitonicSort(a, lo, mid, !dir)
			basicBitonicSort(a, lo+mid, len-mid, dir)
			basicBitonicMerge(a, lo, len, dir)
		}
	}
}

func basicGreatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k > 0 && k < n {
		k <<= 1
	}
	return k >> 1
}

func basicBitonicMerge(a []int, lo, len int, dir bool) {
	if len > 1 {
		m := basicGreatestPowerOfTwoLessThan(len)

		for i := lo; i < lo+len-m; i++ {
			if dir == (a[i] > a[i+m]) {
				a[i], a[i+m] = a[i+m], a[i]
			}
		}
		basicBitonicMerge(a, lo, m, dir)
		basicBitonicMerge(a, lo+m, len-m, dir)
	}

}

func basicBitonicSort(a []int, lo, len int, dir bool) {

	if len > 1 {
		mid := len >> 1
		basicBitonicSort(a, lo, mid, !dir)
		basicBitonicSort(a, lo+mid, len-mid, dir)
		basicBitonicMerge(a, lo, len, dir)
	}
}
