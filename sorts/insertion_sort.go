/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func InsertionSort(a []int) {
	insertionSort(a, 0, len(a))
}

func insertionSort(a []int, lo, hi int) {

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
