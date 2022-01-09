/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func BubbleSort(a []int) {
	bubbleSort(a, len(a))
}

func bubbleSort(a []int, len int) {

	for i := 1; i < len; i++ {
		var swapped bool = false

		for j := 0; j < len-i; j++ {
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
