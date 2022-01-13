package sorts

func DualPivotQuickSort(a []int) {
	dualPivotQuickSort(a, 0, len(a))

}

func dualPivotQuickSort(a []int, lo, hi int) {

	if hi-lo < 2 {
		return
	}

	lp, rp := _3wayQuickPartition(a, lo, hi-1)
	dualPivotQuickSort(a, lo, lp)
	dualPivotQuickSort(a, lp+1, rp)
	dualPivotQuickSort(a, rp+1, hi)

}

func _3wayQuickPartition(a []int, lo, hi int) (int, int) {

	if a[lo] > a[hi] {
		a[lo], a[hi] = a[hi], a[lo]
	}

	l := lo + 1
	k := lo + 1
	g := hi - 1

	leftPivot := a[lo]
	rightPivot := a[hi]

	for k <= g {
		if a[k] < leftPivot {
			a[k], a[l] = a[l], a[k]
			l++
		} else if a[k] >= rightPivot {
			for rightPivot < a[g] && k < g {
				g--
			}
			a[k], a[g] = a[g], a[k]
			g--

			if a[k] < leftPivot {
				a[k], a[l] = a[l], a[k]
				l++
			}
		}
		k++
	}
	l--
	g++

	a[lo], a[l] = a[l], a[lo]
	a[hi], a[g] = a[g], a[hi]

	return l, g
}
