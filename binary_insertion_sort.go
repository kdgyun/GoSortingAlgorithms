package main

func getAscending(a []int, lo, hi int) int {
	limit := lo + 1
	if limit == hi {
		return 1
	}

	if a[lo] <= a[limit] {
		for limit < hi && a[limit-1] <= a[limit] {
			limit++
		}
	} else {
		for limit < hi && a[limit-1] > a[limit] {
			limit++
		}
		reversing(a, lo, limit)
	}
	return limit - lo
}

func reversing(a []int, lo, hi int) {
	hi--
	for lo < hi {
		a[lo], a[hi] = a[hi], a[lo]
		lo++
		hi--
	}
}

func BinarySort(a []int) {
	if len(a) < 2 {
		return
	}
	incLength := getAscending(a, 0, len(a))
	binarySort(a, 0, len(a), incLength)
}

func binarySort(a []int, lo, hi, start int) {

	if lo == start {
		start++
	}

	for ; start < hi; start++ {
		target := a[start]

		locate := binarySearch(a, target, lo, start)
		j := start - 1

		for j >= locate {
			a[j+1] = a[j]
			j--
		}
		a[locate] = target
	}
}

func binarySearch(a []int, key, lo, hi int) int {

	for lo < hi {
		mid := lo + ((hi - lo) >> 1)

		if key < a[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}

	}
	return lo
}
