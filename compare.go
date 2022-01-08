/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

import "fmt"

func Equal(verify, b []int) (bool, string) {
	if len(verify) != len(b) {
		fmt.Errorf("\tdiff : length\t")
		return false, "\tdiff : length\t"
	}
	for i, v := range verify {
		if v != b[i] {
			s := fmt.Sprintf("\tdiff : elements\tstart idx %d [...(correct v : %d, incorrect v : %d) ...]",
				i, verify[i], b[i])
			fmt.Errorf(s)
			return false, s
		}
	}
	return true, ""
}
