/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func OddEvenMergeSort(a []int) {
	oddEvenMergeSort(a, 0, len(a))
}

func oddEvenMergeSort(a []int, lo, hi int) {
	if hi-lo > 1 {
		mid := (hi + lo) >> 1
		oddEvenMergeSort(a, lo, mid)
		oddEvenMergeSort(a, mid, hi)
		oddEvenMerge(a, lo, hi, 1)
	}
}

func oddEvenMerge(a []int, lo, hi, dist int) {
	subDist := dist << 1
	if subDist < (hi - lo) {
		oddEvenMerge(a, lo, hi, subDist)
		oddEvenMerge(a, lo+dist, hi, subDist)
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
