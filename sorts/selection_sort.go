/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func SelectionSort(a []int) {
	selectionSort(a, len(a))
}

func selectionSort(a []int, len int) {
	for i := 0; i < len-1; i++ {
		minIdx := i

		for j := i + 1; j < len; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[minIdx], a[i] = a[i], a[minIdx]
	}
}
