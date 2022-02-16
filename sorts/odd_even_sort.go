package sorts

func OddEvenSort(a []int) {
	oddEvenSort(a, 0, len(a))
}

func oddEvenSort(a []int, lo, hi int) {
	sorted := false
	for !sorted {
		sorted = true

		for i := lo; i < hi-1; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}

		for i := lo + 1; i < hi-1; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}

	}
}
