package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MergeSort(a []int) {
	mergeSort(a, len(a))
}

func mergeSort(a []int, len int) {
	bottomUpMergeSort(a, 0, len-1)
}

func bottomUpMergeSort(a []int, left int, right int) {
	temp := make([]int, right+1)
	for size := 1; size <= right; size += size {

		for l := 0; l <= right-size; l += (2 * size) {
			high := min(l+(2*size)-1, right)
			merge(a, temp, l, l+size-1, high)
		}
	}
}

func merge(origin, temp []int, left, mid, right int) {
	l := left
	r := mid + 1
	idx := left
	for l <= mid && r <= right {
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
	if l > mid {
		for r <= right {
			temp[idx] = origin[r]
			idx++
			r++
		}
	} else {
		for l <= mid {
			temp[idx] = origin[l]
			idx++
			l++
		}
	}

	copy(origin[left:], temp[left:right+1])
}
