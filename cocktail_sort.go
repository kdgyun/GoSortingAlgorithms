/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

func CocktailSort(a []int) {
	cocktailSort(a, len(a))
}

func cocktailSort(a []int, len int) {
	lo := 0
	hi := len
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
