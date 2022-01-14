/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func BubbleSort(a []int) {
	bubbleSort(a, 0, len(a))
}

func bubbleSort(a []int, lo, hi int) {

	for i := lo + 1; i < hi; i++ {
		var swapped bool = false

		for j := lo; j < hi-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j] // swap
				swapped = true
			}
		}
		if !swapped { // [optimizing] if array is ordered, exit this method .
			break
		}
	}
}
