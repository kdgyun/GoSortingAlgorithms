/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func HeapSort(a []int) {
	heapSort(a, len(a))
}

func heapSort(a []int, len int) {
	if len < 2 {
		return
	}
	last := len - 1
	parent := ((last - 1) >> 1)

	for parent >= 0 {
		heapify(a, parent, last)
		parent--
	}

	end := len - 1
	for end > 0 {
		a[0], a[end] = a[end], a[0]
		end--
		heapify(a, 0, end)
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
