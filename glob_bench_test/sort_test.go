package glob_bench_test

import (
	"reflect"
	"testing"

	"github.com/kdgyun/GoSortingAlgorithms/simplytest"
	"github.com/kdgyun/GoSortingAlgorithms/sorts"
)

func TestSort(t *testing.T) {
	size := 1
	for i := 0; i < 6; i++ {
		origin, verify := simplytest.RANDOM_INT(size)

		test := make([]int, len(origin))
		copy(test, origin)
		t.Run("bubble sort", func(t *testing.T) { sorts.BubbleSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [bubble sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("cocktail sort", func(t *testing.T) { sorts.CocktailSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [cocktail sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("insertion sort", func(t *testing.T) { sorts.InsertionSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [insertion sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("selection sort", func(t *testing.T) { sorts.SelectionSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [selection sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("shell sort", func(t *testing.T) { sorts.ShellSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [shell sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("bottom-up merge sort", func(t *testing.T) { sorts.BottomUpMergeSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [bottom-up merge]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("top-down merge sort", func(t *testing.T) { sorts.TopDownMergeSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [top-down merge]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel merge sort", func(t *testing.T) { sorts.ParallelMergeSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel merge]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("heap sort", func(t *testing.T) { sorts.HeapSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [heap sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("left-pivot quick sort", func(t *testing.T) { sorts.QuickSortLP(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [left-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("middle-pivot quick sort", func(t *testing.T) { sorts.QuickSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [middle-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("right-pivot quick sort", func(t *testing.T) { sorts.QuickSortRP(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [right-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel left-pivot quick sort", func(t *testing.T) { sorts.ParallelQuickSortLP(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel left-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel middle-pivot quick sort", func(t *testing.T) { sorts.ParallelQuickSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel middle-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel right-pivot quick sort", func(t *testing.T) { sorts.ParallelQuickSortRP(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel right-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("dual-pivot quick sort", func(t *testing.T) { sorts.DualPivotQuickSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [dual-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel dual-pivot quick sort", func(t *testing.T) { sorts.ParallelDualPivotQuickSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel dual-pivot quick sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("binary insertion sort", func(t *testing.T) { sorts.BinarySort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [binary insertion sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("tim sort", func(t *testing.T) { sorts.TimSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [tim sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("bitonic sort", func(t *testing.T) { sorts.BitonicSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [bitonic sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel bitonic sort", func(t *testing.T) { sorts.ParallelBitonicSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel bitonic sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("intro sort", func(t *testing.T) { sorts.IntroSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [intro sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("parallel intro sort", func(t *testing.T) { sorts.ParallelIntroSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [parallel intro sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("cycle sort", func(t *testing.T) { sorts.CycleSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [cycle sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("odd-even sort", func(t *testing.T) { sorts.OddEvenSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [odd-even sort]")
		}

		test = make([]int, len(origin))
		copy(test, origin)
		t.Run("odd-even merge sort", func(t *testing.T) { sorts.OddEvenMergeSort(test) })

		if !reflect.DeepEqual(test, verify) {
			t.Error("Wrong result [odd-even sort]")
		}
		size <<= 2
	}
}
