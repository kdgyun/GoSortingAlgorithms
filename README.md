# GoSortingAlgorithms

<br />

 Various Sorting Algorithms with golang :octocat:
   

<br /><br />    

## :beginner: Installation   

Once you have installed Go, run the go get command to install the GoSoringAlgorithms package:   

```shell
$ go get github.com/kdgyun/GoSortingAlgorithms
```   


<br /><br />   

## :book: Usage   

it's simple!

```go

package main

import (
	"github.com/kdgyun/GoSortingAlgorithms/sorts"

	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
)

// generate random data
func SliceBuilder(len int) []int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())

	var slice []int
	for i := 0; i < len; i++ {
		slice = append(slice, rand.Int())
	}
	return slice
}

func main() {
	slice := SliceBuilder(1000000)
 
	sorts.QuickSort(slice) // sorts.____(slice []int)

	isSorted := sort.SliceIsSorted(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Print(isSorted)
}

```



<br /><br />   

## :test_tube: Simply Test   

It's easy to get test of all sorting algorithms with simplytest. 

```go

package main

import (
	"github.com/kdgyun/GoSortingAlgorithms/simplytest"
)

func main() {
	simplytest.RunTest()
}


```   

<br />   
<br />   

<b>Example of Output</b>   

```

+==================================================================================+
|                                   Random Test                                    |
+==================================================================================+

...

[array length : 100000]
make arrays...
runing bubble sort...
runing cocktail sort...

...

runing intro sort...
runing parallel intro sort...
runing cycle sort...

+-------------------------------------------------------------------------------------------------+
|                                name |                      ns |                 ms |     verify |      (err mag) 
|-------------------------------------------------------------------------------------------------|
|                         bubble sort |          13016202738 ns |           13016 ms |       true |
|-------------------------------------------------------------------------------------------------|
|                       cocktail sort |           8922656474 ns |            8922 ms |       true |
|-------------------------------------------------------------------------------------------------|
|                            tim sort |             11419013 ns |              11 ms |       true |
|-------------------------------------------------------------------------------------------------|
|                        bitonic sort |             32333072 ns |              32 ms |       true |
|-------------------------------------------------------------------------------------------------|

...

|-------------------------------------------------------------------------------------------------|
|                          intro sort |              6665792 ns |               6 ms |       true |
|-------------------------------------------------------------------------------------------------|
|                 parallel intro sort |              2537508 ns |               2 ms |       true |
|-------------------------------------------------------------------------------------------------|
|                          cycle sort |          20209957747 ns |           20209 ms |       true |
+-------------------------------------------------------------------------------------------------+

```

<br />   
<br />   



### Option   

There is an ```option.go``` file so that users can easily adjust the test, and you can adjust three major options in that file.   


1. Select the sorting algorithm.
   If you want to test only a specific sorting algorithm, find the sorting name you want to change and change the variable to true or false. (True : enable, false : disable)  
   > ex) <b> SHELL_SORT  Activate = false </b>   

2. Change the length of the slice to be tested.
	To change the length of the slice to be tested, configure the slice of the variable 'lengths' to the desired value.  
   > ex) <b> ASCENDING_TEST Activate = false </b>   


3. You can activate or disable the data type of slice to be tested. Basically, all slices consisting of ascending, descending, and random data are tested. However, if you want to disable a specific data type, change the variable to false.   
   > ex) <b> var lengths = [...]int{35, 50000, 100000} </b>




<br />   
<br />   



## :mag_right: SUMMARY

Algorithms covered so far:

| name | function name | 
| --- | :-: | 
| [Bubble Sort](#bubble-sort) | BubbleSort |
| [Cocktail Sort](#cocktail-sort) | CocktailSort |
| [Insertion Sort](#insertion-sort) | InsertionSort |
| [Selection Sort](#selection-sort) | SelectionSort |
| [Shell Sort](#shell-sort) | ShellSort |
| [Merge Sort (bottom-up)](#merge-sort) | BottomUpMergeSort |
| [Merge Sort (top-down)](#merge-sort) | TopDownMergeSort |
| [Merge Sort (parallel)](#merge-sort) | ParallelMergeSort |
| [Heap Sort](#heap-sort) | HeapSort |
| [Quick Sort (left-pivot)](#quick-sort) | QuickSortLP |
| [Quick Sort (middle-pivot)](#quick-sort)  | QuickSort |
| [Quick Sort (left-pivot)](#quick-sort)  | QuickSortRP |
| [Quick Sort (left-pivot with parallel)](#quick-sort)  | ParallelQuickSortLP |
| [Quick Sort (middle-pivot with parallel)](#quick-sort)  | ParallelQuickSort |
| [Quick Sort (left-pivot with parallel)](#quick-sort)  | ParallelQuickSortRP |
| [Dual-pivot Quick Sort](#dual-pivot-quick-sort)  | DualPivotQuickSort |
| [Dual-pivot Quick Sort (parallel)](#dual-pivot-quick-sort)  | ParallelDualPivotQuickSort |
| [Binaray Insertion Sort](#binary-insertion-sort)  | BinarySort |
| [Tim Sort](#tim-sort)  | TimSort |
| [Bitonic Sort](#bitonic-sort)  | BitonicSort |
| [Bitonic Sort (parallel)](#bitonic-sort) | ParallelBitonicSort |
| [Intro Sort](#intro-sort) | IntroSort |
| [Intro Sort (parallel)](#intro-sort) | ParallelIntroSort |
| [Cycle Sort](#cycle-sort) | CycleSort |


<br />
<br />

## Bubble Sort

<br />
Bubble sort repeatedly compares and swaps adjacent elements through the list.

This implementation is optimized. so, if the slice(array) is sorted, exit the bubble sort method. If you don't want to optimize, delete the swapped variable in the bubbleSort method.<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) or ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Cocktail Sort

<br />
Cocktail Sort is a variation of Bubble sort.
it is also known as Cocktail shaker sort, bidirectional bubble sort, cocktail sort, shaker sort.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Insertion Sort

<br />
Insertion sort is a method of finding element from the list through the iteration to place in the correct position.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Selection Sort

<br />
Selection sort is like finding minimum element from the unsorted part through each iteration to place it in front position. 
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Shell Sort

<br />
Shell sort is an extended version of insertion sort. This algorithm first sorts the elements that are far away from each other, then it subsequently reduces the gap between them.

<br />
<br />

\*The applied gap sequence is the [**Ciura sequence.**](https://oeis.org/A102549)
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) or ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | depends on gap sequence | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) or ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Merge Sort

<br />

Merge sort works on the basis of **Divide and Conquer algorithm.**  

Conceptually, a merge sort works as follows:

1. Divide the unsorted list into n sublists, each containing one element (a list of one element is considered sorted).
2. Repeatedly merge sublists to produce new sorted sublists until there is only one sublist remaining. This will be the sorted list.   

<br />
<br />

it supports 3 types of algorithms
- **top-down**   
Top-down merge sort algorithm that recursively splits the list into sublists until sublist size is 1, then merges those sublists to produce a sorted list.
- **bottom-up**   
Bottom-up merge sort algorithm which treats the slice(array) as an slice of n sublists of size 1, and iteratively merges sub-lists back and forth between two buffers
- **parallel**   
Through parallelization of recursive calls, the parallel merge sorting algorithm recursively splits the list into sublists until the sublist size is smaller than the threshold, and then merges the sublists to create a sorted list.
And sub-lists smaller than the threshold use a bottom-up list to produce a sorted list.

<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | No | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(n)](https://latex.codecogs.com/svg.image?O(n)) |



<br />
<br />

## Heap Sort

<br />
Heap sort is to eliminate the elements one by one from the heap part of the list, and then insert them into the sorted part of the list.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Quick Sort

<br />
Quick Sort is based on Divide-and-conquer algorithm where it selects a pivot element from the array and then partition the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.   

<br />
<br />

it supports 3 types of algorithms

- **left-pivot (+parallel)**   
Left-pivot Quick sort is implemented with left element selected as the pivot
- **middle-pivot (+parallel)**   
Middle-pivot Quick sort is implemented with middle element selected as the pivot
- **right-pivot (+parallel)**   
right-pivot Quick sort is implemented with right element selected as the pivot   

Through parallelization of recursive calls, each parallel quick sorting algorithms recursively splits the list into sublists until the sublist size is smaller than the threshold, and then sorts the sublists that smaller than the threshold by use a basic quick sort to create a sorted list.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(log_n)](https://latex.codecogs.com/svg.image?O(\log&space;n)) |



<br />
<br />

## Dual-Pivot Quick Sort

<br />
Dual-Pivot Quick Sort is based on Divide-and-conquer algorithm of quick sort, but there is a difference in selecting two elements (pivot) in the list to sort.

Pick two elements(pivot) from the array to be sorted, partition the remaining elements into those less than the lesser pivot, those between the pivots, and those greater than the greater pivot, and recursively sort these partitions.   

it is also known as <b>3-way quick sort</b>.   

<br />
<br />

it supports 2 types of algorithms

- **non-parallel**   
Pick two elements(pivot) from the array to be sorted, partition the remaining elements into those less than the lesser pivot, those between the pivots, and those greater than the greater pivot, and recursively sort these partitions.
- **parallel**   
Through parallelization of recursive calls, each parallel dual-pivot quick sorting algorithms recursively splits the list into sublists until the sublist size is smaller than the threshold, and then sorts the sublists that smaller than the threshold by use a basic dual-pivot quick sort to create a sorted list.



<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(log_n)](https://latex.codecogs.com/svg.image?O(\log&space;n)) |



<br />
<br />

## Binary Insertion Sort

<br />
Binary Insertion Sort is an extended version of insertion sort. it also known as Binary Sort.

Binary Insertion Sort uses binary search to find the proper location to insert the selected item at each iteration. 
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Tim Sort

<br />

Timsort is a  **hybrid stable sorting algorithm**, derived from merge sort and insertion sort(or binary insertion sort), designed to perform well on many kinds of real-world data. It was implemented by Tim Peters in 2002 for use in the Python programming language. The algorithm finds subsequences of the data that are already ordered (runs) and uses them to sort the remainder more efficiently.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(n)](https://latex.codecogs.com/svg.image?O(n)) |



<br />
<br />

## Bitonic Sort

<br />
Bitonic sort is a comparison-based sorting algorithm that can be run in parallel. It focuses on converting a random sequence of numbers into a bitonic sequence, one that monotonically increases, then decreases.

<br />
<br />

it supports 2 types of algorithms
- **non-parallel**   
non-parallel bitonic sort that recursively splits the list into sublists until sublist size is 1, then merges those sublists according to the Bitonic Sequencing Rule to produce a sorted list. 
- **parallel**   
Through parallelization of recursive calls, the parallel bitonic sorting algorithm recursively splits the list into sublists until the sublist size is smaller than the threshold, and then switches to non-parallelization to split and merge the sublists according to the Bitonic Sequencing Rule to create a sorted list.

<br />

### COMPLEXITY


| Type | Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| non-parallel | ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | Yes | No | total : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)), auxiliary : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) |
| parallel | ![O(log^2_n)](https://latex.codecogs.com/svg.image?O(\log^2&space;n)) | ![O(log^2_n)](https://latex.codecogs.com/svg.image?O(\log^2&space;n)) | ![O(log^2_n)](https://latex.codecogs.com/svg.image?O(\log^2&space;n)) | Yes | No | total : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)), auxiliary : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) |



<br />
<br />

## Intro Sort

<br />   

Introsort is a hybrid sorting algorithm that provides both fast average performance and (asymptotically) optimal worst-case performance. It begins with quicksort, it switches to heapsort when the recursion depth exceeds a level based on
(maximum depth:
<span> ![2ceil(log2_n)](https://latex.codecogs.com/svg.image?2&space;&space;\left&space;\lfloor&space;\log_2(n)&space;\right&space;\rfloor) </span>
of)
the number of elements being sorted and it switches to insertion sort when the number of elements is below some threshold(16).   
   
<br />
<br />


it supports 2 types of algorithms

- **non-parallel**   
Pick a element(pivot) when uses quick sort from the array to be sorted, partition the other elements into two sub-arrays according to whether they are less than or greater than the pivot, and recursively sort these partitions until the recursion depth exceeds a level based on
(maximum depth:
<span> ![2ceil(log2_n)](https://latex.codecogs.com/svg.image?2&space;&space;\left&space;\lfloor&space;\log_2(n)&space;\right&space;\rfloor) </span>
of)
the number of elements.
- **parallel**   
Through parallelization of recursive calls, each parallel dual-pivot quick sorting algorithms recursively splits the list into sublists until the sublist size is smaller than the threshold, and then sorts the sublists that smaller than the threshold by use a basic dual-pivot quick sort to create a sorted list.

<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) |  ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(log_n)](https://latex.codecogs.com/svg.image?O(\log&space;n)) |



<br />
<br />

## Cycle Sort

<br />
Cycle sort is an in-place sorting Algorithm, unstable sorting algorithm, a comparison sort that is theoretically optimal in terms of the total number of writes to the original array.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |
