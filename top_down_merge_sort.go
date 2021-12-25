package main

/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

func TopDownMergeSort(a []int) {
	temp := make([]int, len(a))
	topDownMergeSort(a, temp, 0, len(a))
}

func topDownMergeSort(a, temp []int, left int, right int) {

	if right-left < 2 {
		return
	}
	mid := (left + right) / 2
	topDownMergeSort(a, temp, left, mid)
	topDownMergeSort(a, temp, mid, right)
	mergeForTopDown(a, temp, left, mid, right)
}

func mergeForTopDown(origin, temp []int, left, mid, right int) {
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
