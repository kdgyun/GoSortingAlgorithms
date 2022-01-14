/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package sorts

func CocktailSort(a []int) {
	cocktailSort(a, 0, len(a))
}

func cocktailSort(a []int, left, right int) {
	lo := left
	hi := right
	var swapped bool = true

	for swapped {
		swapped = false

		for i := lo; i < hi-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		swapped = false

		hi--
		for i := hi - 1; i >= lo; i-- {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		lo++

	}
}
