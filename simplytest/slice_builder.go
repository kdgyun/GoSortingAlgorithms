/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package simplytest

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"

	"github.com/logrusorgru/aurora/v3"
)

func RANDOM_INT(len int) ([]int, []int) {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	var origin, verify []int
	fmt.Println(aurora.BrightCyan("[array/slice length : " + fmt.Sprint(len) + "]"))

	fmt.Println("make slice...")
	for i := 0; i < len; i++ {
		v := rand.Int()

		verify = append(verify, v)
		origin = append(origin, v)
	}
	sort.Ints(verify)
	return origin, verify
}

func DESCENDING_INT(len int) ([]int, []int) {
	var origin, verify []int
	fmt.Println(aurora.BrightCyan("[array/slice length : " + fmt.Sprint(len) + "]"))

	fmt.Println("make slice...")
	for i := 0; i < len; i++ {

		verify = append(verify, len-i)
		origin = append(origin, len-i)
	}
	sort.Ints(verify)
	return origin, verify
}

func ASCENDING_INT(len int) ([]int, []int) {
	var origin, verify []int
	fmt.Println(aurora.BrightCyan("[array/slice length : " + fmt.Sprint(len) + "]"))

	fmt.Println("make slice...")
	for i := 0; i < len; i++ {

		verify = append(verify, i)
		origin = append(origin, i)
	}
	sort.Ints(verify)
	return origin, verify
}
