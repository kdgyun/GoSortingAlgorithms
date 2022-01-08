/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

func BottomUpMergeSort(a []int) {
	bottomUpMergeSort(a, 0, len(a))
}

func bottomUpMergeSort(a []int, left int, right int) {
	temp := make([]int, right)
	for size := 1; size < right; size += size {

		for l := 0; l < right-size; l += (2 * size) {
			var high int
			mid := l + size

			if l+(2*size) < right {
				high = l + (2 * size)
			} else {
				high = right
			}

			mergeForBottomUp(a, temp, l, mid, high)
		}
	}
}

func mergeForBottomUp(origin, temp []int, left, mid, right int) {
	l := left
	r := mid
	idx := left
	for l < mid && r < right {
		if origin[l] <= origin[r] {
			temp[idx] = origin[l]
			idx++
			l++
		} else {
			temp[idx] = origin[r]
			idx++
			r++
		}
	}
	if l == mid {
		for r < right {
			temp[idx] = origin[r]
			idx++
			r++
		}
	} else {
		for l < mid {
			temp[idx] = origin[l]
			idx++
			l++
		}
	}

	copy(origin[left:], temp[left:right])
}
