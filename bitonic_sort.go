/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

func bitonicMerge(a []int, lo, len int, dir bool) {
	if len > 1 {
		m := greatestPowerOfTwoLessThan(len)

		for i := lo; i < lo+len-m; i++ {
			if dir == (a[i] > a[i+m]) {
				a[i], a[i+m] = a[i+m], a[i]
			}
		}
		bitonicMerge(a, lo, m, dir)
		bitonicMerge(a, lo+m, len-m, dir)
	}

}

func BitonicSort(a []int) {
	bitonicSort(a, 0, len(a), true)
}

func bitonicSort(a []int, lo, len int, dir bool) {

	if len > 1 {
		mid := len >> 1
		bitonicSort(a, lo, mid, !dir)
		bitonicSort(a, lo+mid, len-mid, dir)
		bitonicMerge(a, lo, len, dir)
	}
}

func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k > 0 && k < n {
		k <<= 1
	}
	return k >> 1
}
func BottomUpBitonicSort(a []int) {

}

func bottomUpBitonicSort(a []int, len int) {

	for blockSize := 2; blockSize <= len; blockSize <<= 1 {
		for subPart := blockSize / 2; subPart > 0; subPart >>= 1 {

			for i := 0; i < len; i++ {

			}
		}
	}
}
