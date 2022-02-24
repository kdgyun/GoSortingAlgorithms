/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package simplytest

type Activate bool

// Section 1.
// sorting algorithm test list
const (
	BUBBLE_SORT                      Activate = true
	COCKTAIL_SORT                    Activate = true
	INSERTION_SORT                   Activate = true
	SELECTION_SORT                   Activate = true
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
	PARALLEL_DUAL_PIVOT_QUICK_SORT   Activate = true
	BINARY_INSERTION_SORT            Activate = true
	TIM_SORT                         Activate = true
	BITONIC_SORT                     Activate = true
	PARALLEL_BITONIC_SORT            Activate = true
	INTRO_SORT                       Activate = true
	PARALLEL_INTRO_SORT              Activate = true
	CYCLE_SORT                       Activate = true
	ODDEVEN_SORT                     Activate = true
	ODDEVEN_MERGE_SORT               Activate = true // The length n of the array must be a power of 2
	PARALLEL_ODDEVEN_MERGE_SORT      Activate = true // The length n of the array must be a power of 2
)

// Section 2.
// The type of slice to be tested.
const (
	ASCENDING_TEST  Activate = true
	RANDOM_TEST     Activate = true
	DESCENDING_TEST Activate = true
)

// Section 3.
// lengths for test
var lengths = [...]int{1 << 4, 1 << 7, 1 << 10, 1 << 13, 1 << 16}
