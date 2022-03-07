/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func CombSort(a []int) {
	combSort(a, 0, len(a))
}

func combSort(a []int, lo, hi int) {
	gap := hi - lo
	shrink := 1.3
	sorted := false

	for !sorted {
		gap = int(float64(gap) / (shrink))
		if gap <= 1 {
			sorted = true
			gap = 1
		}

		for i := lo; i < hi-gap; i++ {
			if a[i] > a[i+gap] {
				a[i], a[i+gap] = a[i+gap], a[i]
				sorted = false
			}
		}
	}
}
