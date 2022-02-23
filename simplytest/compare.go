/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package simplytest

import "fmt"

func Equal(verify, b []int) (bool, string) {
	if len(verify) != len(b) {
		fmt.Errorf("\tdiff : length\t")
		return false, "\tdiff : length\t"
	}
	for i, v := range verify {
		if v != b[i] {
			s := fmt.Sprintf("\t1elements not sorted",
				i, verify[i], b[i])
			fmt.Errorf(s)
			return false, s
		}
	}
	return true, ""
}
