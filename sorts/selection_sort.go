/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func SelectionSort(a []int) {
	selectionSort(a, 0, len(a))
}

func selectionSort(a []int, lo, hi int) {
	for i := lo; i < hi-1; i++ {
		minIdx := i

		for j := i + 1; j < hi; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[minIdx], a[i] = a[i], a[minIdx]
	}
}
