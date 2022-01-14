/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func HeapSort(a []int) {
	heapSort(a, 0, len(a))
}

func heapSort(a []int, lo, hi int) {
	if hi-lo < 2 {
		return
	}
	last := hi - 1
	parent := ((last - 1) >> 1)

	for parent >= 0 {
		heapify(a, parent, last)
		parent--
	}

	end := hi - 1
	for end > lo {
		a[lo], a[end] = a[end], a[lo]
		end--
		heapify(a, lo, end)
	}
}

func heapify(a []int, cur, last int) {
	var left, right, large int

	for ((cur << 1) + 1) <= last {
		left = (cur << 1) + 1
		right = (cur << 1) + 2
		large = cur

		if a[left] > a[large] {
			large = left
		}

		if right <= last && a[right] > a[large] {
			large = right
		}

		if large != cur {
			a[cur], a[large] = a[large], a[cur]
			cur = large
		} else {
			return
		}
	}
}
