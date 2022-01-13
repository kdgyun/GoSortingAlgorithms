/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

type Activate bool

const (
	BUBBLE_SORT                      Activate = false
	COCKTAIL_SORT                    Activate = false
	INSERTION_SORT                   Activate = false
	SELECTION_SORT                   Activate = false
	SHELL_SORT                       Activate = true
	BOTTOM_UP_MERGE_SORT             Activate = true
	TOP_DOWN_MERGE_SORT              Activate = true
	PARALLEL_MERGE_SORT              Activate = true
	HEAP_SORT                        Activate = true
	LEFT_PIVOT_QUICK_SORT            Activate = true
	MIDDLE_PIVOT_QUICK_SORT          Activate = true
	RIGHT_PIVOT_QUICK_SORT           Activate = true
	PARALLEL_LEFT_PIVOT_QUICK_SORT   Activate = true
	PARALLEL_MIDDLE_PIVOT_QUICK_SORT Activate = true
	PARALLEL_RIGHT_PIVOT_QUICK_SORT  Activate = true
	DUAL_PIVOT_QUICK_SORT            Activate = true
	BINARY_INSERTION_SORT            Activate = false
	TIM_SORT                         Activate = true
	BITONIC_SORT                     Activate = true
	PARALLEL_BITONIC_SORT            Activate = true
)

var lengths = [...]int{10, 100, 10000, 100000, 1000000, 10000000} // lengths for test

func main() {
	// AscendingTest()
	// DescendingTest()
	RandomTest()
}
