package main

type Activate bool

const (
	BUBBLE_SORT                    Activate = true
	INSERTION_SORT                 Activate = true
	SELECTION_SORT                 Activate = true
	SHELL_SORT                     Activate = true
	BOTTOM_UP_MERGE_SORT           Activate = true
	TOP_DOWN_MERGE_SORT            Activate = true
	PARALLEL_MERGE_SORT            Activate = true
	HEAP_SORT                      Activate = true
	LEFT_PIVOT_QUICK_SORT          Activate = true
	MIDDLE_PIVOT_QUICK_SORT        Activate = true
	RIGHT_PIVOT_QUICK_SORT         Activate = true
	PARALLEL_LEFT_PIVOT_QUICK_SORT Activate = true
)

var length = [...]int{10, 100, 1000, 10000, 100000, 1000000} // length of slice to test

func main() {
	AscendingTest()
	DescendingTest()
	RandomTest()
}
