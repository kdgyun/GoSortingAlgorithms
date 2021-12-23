/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

/*
 Gap sequence is based on the Ciura sequence.
 The number after 1750 is obtained by multiplying by 2.25.
 link : https://en.wikipedia.org/wiki/Shellsort#Gap_sequences
*/
var Gap = [...]int{1, 4, 10, 23, 57, 132, 301, 701, 1750, 3937, 8858, 19930, 44842, 100894, 227011, 510774, 1149241, 2585792, 5818032, 13090572, 29453787, 66271020, 149109795, 335497038, 754868335, 1698453753}

func getGap(len int) int {
	idx := 0
	threshold := int(float32(len) / 2.25)

	for Gap[idx] < threshold {
		idx++
	}
	return idx
}

func ShellSort(a []int) {
	shellSort(a, len(a))
}

func shellSort(a []int, len int) {

	pos := getGap(len)

	for pos >= 0 {
		step := Gap[pos]
		pos--

		for i := step; i < len; i++ {
			target := a[i]
			j := i

			for ; j >= step && target < a[j-step]; j -= step {
				a[j] = a[j-step]
			}
			a[j] = target
		}
	}
}
