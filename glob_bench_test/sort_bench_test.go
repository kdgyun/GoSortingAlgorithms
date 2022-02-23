package glob_bench_test

import (
	"math"
	"math/big"
	"math/rand"

	crand "crypto/rand"
	"testing"

	"github.com/kdgyun/GoSortingAlgorithms/sorts"
)

var lengthforbench = 1 << 16

func sliceBuilder(len int) []int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	var a []int

	for i := 0; i < len; i++ {
		v := rand.Int()
		a = append(a, v)
	}

	return a
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.BubbleSort(test)
		b.StopTimer()
	}
}

func BenchmarkCocktailSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.CocktailSort(test)
		b.StopTimer()
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.InsertionSort(test)
		b.StopTimer()
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.SelectionSort(test)
		b.StopTimer()
	}
}

func BenchmarkShellSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ShellSort(test)
		b.StopTimer()
	}
}

func BenchmarkBottomUpMergeSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.BottomUpMergeSort(test)
		b.StopTimer()
	}
}

func BenchmarkTopDownMergeSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.TopDownMergeSort(test)
		b.StopTimer()
	}
}

func BenchmarkParallelMergeSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelMergeSort(test)
		b.StopTimer()
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.HeapSort(test)
		b.StopTimer()
	}
}

func BenchmarkQuickSortLP(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.QuickSortLP(test)
		b.StopTimer()
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.QuickSort(test)
		b.StopTimer()
	}
}

func BenchmarkQuickSortRP(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.QuickSortRP(test)
		b.StopTimer()
	}
}

func BenchmarkParallelQuickSortLP(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelQuickSortLP(test)
		b.StopTimer()
	}
}

func BenchmarkParallelQuickSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelQuickSort(test)
		b.StopTimer()
	}
}

func BenchmarkParallelQuickSortRP(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelQuickSortRP(test)
		b.StopTimer()
	}
}

func BenchmarkDualPivotQuickSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.DualPivotQuickSort(test)
		b.StopTimer()
	}
}

func BenchmarkParallelDualPivotQuickSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelDualPivotQuickSort(test)
		b.StopTimer()
	}
}

func BenchmarkBinarySort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.BinarySort(test)
		b.StopTimer()
	}
}

func BenchmarkTimSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.TimSort(test)
		b.StopTimer()
	}
}

func BenchmarkBitonicSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.BitonicSort(test)
		b.StopTimer()
	}
}

func BenchmarkParallelBitonicSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelBitonicSort(test)
		b.StopTimer()
	}
}

func BenchmarkIntroSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.IntroSort(test)
		b.StopTimer()
	}
}

func BenchmarkParallelIntroSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.ParallelIntroSort(test)
		b.StopTimer()
	}
}

func BenchmarkCycleSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.CycleSort(test)
		b.StopTimer()
	}
}

func BenchmarkOddEvenSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.OddEvenSort(test)
		b.StopTimer()
	}
}

func BenchmarkOddEvenMergeSort(b *testing.B) {
	b.StopTimer()
	unsorted := sliceBuilder(lengthforbench)
	test := make([]int, len(unsorted))

	for i := 0; i < b.N; i++ {
		copy(test, unsorted)
		b.StartTimer()
		sorts.OddEvenMergeSort(test)
		b.StopTimer()
	}
}
