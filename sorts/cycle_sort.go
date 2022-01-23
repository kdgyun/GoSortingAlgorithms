package sorts

func CycleSort(a []int) {
	cycleSort(a, 0, len(a))
}

func cycleSort(a []int, left, right int) {

	for start := left; start < right; start++ {
		item := a[start]
		pos := start

		for i := start + 1; i < right; i++ {
			if a[i] < item {
				pos++
			}
		}

		if pos == start {
			continue
		}

		for item == a[pos] {
			pos++
		}

		a[pos], item = item, a[pos]

		for pos != start {
			pos = start
			for i := start + 1; i < right; i++ {
				if a[i] < item {
					pos++
				}
			}

			for item == a[pos] {
				pos++
			}
			a[pos], item = item, a[pos]

		}
	}
}
