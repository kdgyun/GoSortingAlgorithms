# GoSortingAlgorithms

<br />

 Various Sorting Algorithms with golang :octocat:
   

<br /><br />


## SUMMARY

Algorithms covered so far:

| name | function name | 
| --- | :-: | 
| Bubble Sort | BubbleSort |
| Cocktail Sort | CocktailSort |
| Insertion Sort | InsertionSort |
| Selection Sort | SelectionSort |
| Shell Sort | ShellSort |
| Merge Sort (bottom-up) | BottomUpMergeSort |
| Merge Sort (top-down) | TopDownMergeSort |
| Merge Sort (parallel) | ParallelMergeSort |
| Heap Sort | HeapSort |
| Quick Sort (left-pivot) | QuickSortLP |
| Quick Sort (middle-pivot) | QuickSort |
| Quick Sort (left-pivot) | QuickSortRP |
| Quick Sort (left-pivot with parallel) | ParallelQuickSortLP |
| Quick Sort (middle-pivot with parallel) | ParallelQuickSort |
| Quick Sort (left-pivot with parallel) | ParallelQuickSortRP |
| Binaray Insertion Sort | BinarySort |
| Tim Sort | Timsort |
| Bitonic Sort | BitonicSort |
| Bitonic Sort (parallel) | ParallelBitonicSort |


<br />
<br />

## Bubble Sort

<br />
Bubble sort repeatedly compares and swaps adjacent elements through the slice(array).

this implementation is optimized. so, if the slice(array) is sorted, exit the bubble sort method. If you don't want to optimize, delete the swapped variable in the bubbleSort method.<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) or ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Cocktail Sort

<br />
Cocktail Sort is a variation of Bubble sort.
it also known as Cocktail shaker sort, bidirectional bubble sort, cocktail sort, shaker sort
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Insertion Sort

<br />
Insertion sort is finding each element from the slice(array) through each iteration to place it in correct position.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |
