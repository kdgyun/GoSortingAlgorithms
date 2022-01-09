/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func InsertionSort(a []int) {
	insertionSort(a, len(a))
}

func insertionSort(a []int, len int) {

	for i := 1; i < len; i++ {
		target := a[i]
		j := i - 1

		for j >= 0 && target < a[j] {
			a[j+1] = a[j]
			j--
		}

		a[j+1] = target
	}
}
