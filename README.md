**한글 문서** | **[English Document](./README-en.md)**

<br>

![](https://img.shields.io/github/go-mod/go-version/kdgyun/GoSortingAlgorithms/main)
![](https://img.shields.io/github/issues/kdgyun/GoSortingAlgorithms) 
![](https://img.shields.io/github/commit-status/kdgyun/GoSortingAlgorithms/main/ee331ab723bf79aa3b98a2c4e143bd07eee5b1f2)
![](https://img.shields.io/github/repo-size/kdgyun/GoSortingAlgorithms)
![](https://github.com/kdgyun/GoSortingAlgorithms/actions/workflows/go.yml/badge.svg?branch=main)

# GoSortingAlgorithms

<br />

 Go 언어로 작성된 다양한 정렬 알고리즘들이 있습니다 :octocat:
   

<br /><br />    

## :beginner: 설치 방법   

Go 언어가 설치 되어 있다면, go get 명령어를 실행하여 GoSoringAlgorithms 패키지를 받아옵니다.    

```shell
$ go get github.com/kdgyun/GoSortingAlgorithms
```   


<br /><br />   

## :book: 사용 설명서   

간단합니다!

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

// 랜덤 정수 슬라이스 생성기
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

## :test_tube: 간단한 테스트 방법   

symplytest 를 통해 쉽게 테스트 할 수 있습니다. 

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

<b>출력 예시</b>   

```

+==================================================================================+
|                                   Random Test                                    |
+==================================================================================+

...

[array length : 100000]
make arrays...
running bubble sort...
running cocktail sort...

...

running intro sort...
running parallel intro sort...
running cycle sort...

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

```option.go```를 통해 사용자가 쉽게 테스트 형식을 조정할 수 있으며, 3가지 주요 옵션을 제공합니다.


1. 정렬 알고리즘 선택
   특정 정렬 알고리즘만 테스트 혹은 제외하고자 한다면 테스트하고자 하는 정렬 알고리즘 이름을 찾아 변수를 true 또는 false로 변경합니다. (True : 테스트 허용, false : 테스트 비허용)  
   > ex) <b> SHELL_SORT  Activate = true  </b>   

2. 테스트 할 슬라이스의 길이 변경.
	테스트할 슬라이스의 길이를 변경하려면 'lengths' 변수의 슬라이스를 원하는 값으로 설정 할 수 있습니다. 아래 예시는 길이 35, 50,000, 100,000 슬라이스에 대해 각각 테스트 하게 됩니다.   
   > ex) <b> var lengths = [...]int{35, 50000, 100000} </b>


3. 슬라이스 유형 변경.   
    기본적으로 오름차순, 내림차순 및 랜덤 데이터로 구성된 각각의 모든 슬라이스가 테스트됩니다. 그러나 특정 데이터 유형을 비활성화하려면 변수를 false로 변경하면 됩니다.   
   > ex) <b> ASCENDING_TEST Activate = false </b>   




<br />   
<br />   



## :mag_right: 개요



Go 언어는 배열과 슬라이스가 구분되어 있습니다. 하지만, 일반적으로 기타 언어에서도 배열이 익숙하기 때문에 해당 알고리즘을 설명 할 때에는 **배열**로 통일하여 설명합니다.

현재 업데이트 된 정렬 알고리즘:

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
| [Odd-Even Sort](#odd-even-sort) | OddEvenSort |
| [Odd-Even Merge Sort](#odd-even-merge-sort) | OddEvenMergeSort |
| [Odd-Even Merge Sort (parallel)](#odd-even-merge-sort) | ParallelOddEvenMergeSort |
| [Comb Sort](#comb-sort) | CombSort |


<br />
<br />

## Bubble Sort

<br />
버블 정렬은 인접한 요소를 반복적으로 비교하고 교환하는 방식입니다.
해당 구현은 이미 배열이 정렬되어있는 경우 정렬을 종료하도록 최적화되어 있습니다.  최적화를 원하지 않는다면, bubbleSort 함수에서 swapped 변수를 삭제 및 해당 조건문을 삭제하면 됩니다.<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) or ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Cocktail Sort

<br />
칵테일 정렬은 버블 정렬을 기반한 변형 된 알고리즘입니다. Cocktail shaker sort(칵테일 셰이커 정렬), bidirectional bubble sort(양방향 버블 정렬), cocktail sort(칵테일 정렬), shaker sort(셰이커 정렬) 이라고도 불립니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Insertion Sort

<br />
삽입 정렬은 각 반복마다 배열의 앞에서부터 차례대로 이미 정렬된 배열 부분과 비교하여 자신의 위치를 찾아 삽입요소를 찾아 올바른 위치에 배치하게 됩니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Selection Sort

<br />
선택 정렬은 정렬되지 않은 부분에서 각 반복을 통해 가장 작은 요소를 찾아 앞 위치에 배치하게 됩니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Shell Sort

<br />
셸 정렬은 삽입 정렬을 확장한 버전입니다. 먼저 일정 간격으로 서로 멀리 떨어져 있는 요소를 삽입 정렬해 나가면서 다음에는 그 사이의 간격을 줄여나가면서 정렬하게 됩니다.

<br />
<br />

\*적용 된 Gap 시퀀스는 다음과 같습니다. [**Ciura sequence**](https://oeis.org/A102549)
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) or ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | depends on gap sequence | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) or ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Merge Sort

<br />

병합 정렬은 **분할 정복 알고리즘**을 기반으로 합니다.

개념적으로 병합 정렬은 다음과 같이 작동합니다:

1. 정렬되지 않은 n개의 요소를 갖는 배열을 재귀적으로 최소 하나 이상의 요소를 포함하는 배열이 되도록 절반으로 나눕니다(요소가 한 개만 있는 배열은 정렬된 것으로 간주됩니다).
2. 분할 된 배열이 모두 합쳐질 때 까지 반복적으로 병합하여 정렬 합니다.   

<br />
<br />

3가지 유형의 병합 정렬을 지원합니다.
- **top-down**   
하향식 병합 정렬 알고리즘은 분할 된 배열의 크기가 1이 될 때까지 배열을 재귀적으로 분할한 다음 해당 하위 배열들을 병합하여 정렬된 목록을 생성합니다.
- **bottom-up**   
상향식 병합 정렬 알고리즘은 크기가 1인 n개의 하위 배열을 두 버퍼 사이에서 앞 뒤로 하위 배열들을 반복적으로 병합합니다.
- **parallel**   
병렬 병합 정렬 알고리즘은 재귀 호출의 병렬화를 통해 하위 배열의 크기가 임계값보다 작아질 때까지 배열을 재귀적으로 분할한 다음 병합하는 하향식 정렬을 수행합니다. 이 때, 임계값보다 작은 하위 배열들은 상향식 정렬을 사용하여 정렬하게 됩니다.   


<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | No | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(n)](https://latex.codecogs.com/svg.image?O(n)) |



<br />
<br />

## Heap Sort

<br />
힙 정렬은 배열을 힙의 원리처럼 배열에서 요소를 하나씩 제거한 다음 배열의 정렬된 부분에 삽입하여 정렬하게 됩니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Quick Sort

<br />
퀵 정렬은 배열에서 특정 피벗을 선택한 다음 피벗보다 작은지 큰지에 따라 다른 요소를 두 개의 하위 배열로 분할하게 되는 분할 정복 알고리즘을 기반으로 합니다. 재귀적으로 이 과정을 거쳐 정렬됩니다.


<br />
<br />

3가지 유형의 퀵 정렬을 지원합니다.

- **left-pivot (+parallel)**   
왼쪽 피벗 퀵 정렬은 왼쪽 요소를 피벗으로 선택하여 정렬합니다.   
- **middle-pivot (+parallel)**   
중간 피벗 퀵 정렬은 중간 요소를 피벗으로 선택하여 정렬합니다.   
- **right-pivot (+parallel)**   
오른쪽 피벗 퀵 정렬은 오른쪽 요소를 피벗으로 선택하여 정렬합니다.   

또한 각 알고리즘은 재귀 호출의 병렬화도 지원합니다. 
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(log_n)](https://latex.codecogs.com/svg.image?O(\log&space;n)) |



<br />
<br />

## Dual-Pivot Quick Sort

<br />
이중 피벗 퀵 정렬은 Quick Sort를 기반으로 하지만, 배열에서 두 개의 요소를 피벗(pivot)으로 선택하여 정렬하는 방식에 차이가 있습니다.
정렬할 배열에서 두 개의 요소(피벗)를 선택하고 나머지 요소들을 작은 피벗(small pivot)보다 작은 요소, 피벗 사이에 있는 요소, 큰 피벗(big pivot)보다 큰 요소로 구분하여 배열을 분할하고 이러한 파티션을 재귀적인 과정을 통해 정렬하게 됩니다.


<br />
<br />

2가지 유형의 이중 피벗 퀵 정렬을 지원합니다.

- **non-parallel**   
정렬할 배열에서 두 개의 요소(피벗)를 선택하고 나머지 요소들을 작은 피벗(small pivot)보다 작은 요소, 피벗 사이에 있는 요소, 큰 피벗(big pivot)보다 큰 요소로 구분하여 배열을 분할하고 이러한 파티션을 재귀적인 과정을 통해 정렬하게 됩니다.
- **parallel**   
재귀 호출의 병렬화를 통해 각각의 정렬 알고리즘은 분할 된 배열의 크기가 임계값보다 작을 때까지 배열을 재귀적으로 분할한 배열을 병합하여 정렬하게 됩니다.



<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(log_n)](https://latex.codecogs.com/svg.image?O(\log&space;n)) |



<br />
<br />

## Binary Insertion Sort

<br />
이진 삽입 정렬은 삽입 정렬을 기반으로 한 확장 버전입니다. 이진 정렬이라고도 합니다.
이진 삽입 정렬은 이진 검색(Binary Search)을 사용하여 각 반복 과정에서 특정 요소를 삽입할 위치를 찾아 삽입하게 됩니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Tim Sort

<br />

Timsort는 다양한 종류의 실제 데이터에서 잘 수행되도록 설계된 알고리즘으로, 병합 정렬과 삽입 정렬(또는 이진 삽입 정렬)에서 파생되어 혼합 된 **안정적인 하이브리드 정렬 알고리즘** 입니다. Python 프로그래밍 언어에서 사용하기 위해 2002년 Tim Peters에 의해 구현되었습니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(n)](https://latex.codecogs.com/svg.image?O(n)) | Yes | yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(n)](https://latex.codecogs.com/svg.image?O(n)) |



<br />
<br />

## Bitonic Sort

<br />
바이토닉 정렬은 병렬로 실행할 수 있는 비교 기반 정렬 알고리즘입니다. 무작위 숫자 시퀀스를 단조 증가했다가 감소하는 비트 시퀀스로 변환하는 데 중점을 둡니다. 병렬 환경에 최적화 되어있으며 GPU 병렬 처리로 사용되기도 하지만, 여기서는 기타 병렬화한 정렬 알고리즘들과 마찬가지로 GO 언어의 동시성 모델로 구현됩니다.    


<br />
<br />

2가지 유형의 알고리즘을 지원합니다.   
- **non-parallel**   
분할 된 배열의 크기가 1이 될 때까지 배열을 재귀적으로 분할한 다음 Bitonic Sequencing Rule에 따라 해당 하위 배열을 병합하여 정렬하게 됩니다.   
- **parallel**   
재귀 호출의 병렬화를 통해 분할 된 배열의 크기가 임계값보다 작아질 때까지 배열을 재귀적으로 분할한 다음 비병렬화로 전환하여 Bitonic Sequencing Rule에 따라 하위 배열을 분할 및 병합하여 생성합니다.   

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

인트로 정렬은 최악의 경우에도 평균적으로 빠른 성능과 및 (점근적으로) 최적화 된 하이브리드 정렬 알고리즘입니다. 퀵정렬로 시작하여 재귀 깊이가 정렬되는 배열의 요소 개수에 따른 깊이 임계값(maximum depth: <span> ![2ceil(log2_n)](https://latex.codecogs.com/svg.image?2&space;&space;\left&space;\lfloor&space;\log_2(n)&space;\right&space;\rfloor) </span> ) 수준을 초과하면 힙 정렬로 전환하고 분할 된 배열의 요소 개수가 길이 임계값(16) 미만이면 삽입 정렬로 전환합니다. 
   
<br />
<br />


2가지 유형의 알고리즘을 지원합니다.  

- **non-parallel**   
정렬할 배열에서 퀵 정렬을 사용할 때 요소(피벗)를 선택하고 선택 된 피벗을 제외한 다른 요소들이 피벗보다 작거나 큰지 여부에 따라 두 개의 하위 배열로 분할하고 재귀 깊이가 특정 임계값(maximum depth: <span> ![2ceil(log2_n)](https://latex.codecogs.com/svg.image?2&space;&space;\left&space;\lfloor&space;\log_2(n)&space;\right&space;\rfloor) </span>)을 넘기 전까지 이러한 과정을 재귀적으로 거칩니다.

- **parallel**   
재귀 호출의 병렬화를 통해 각 병렬화 된 이중 피벗 퀵 정렬 알고리즘은 분할 된 배열의 크기가 임계값보다 작아지기 전까지 배열을 재귀적으로 분할하여 정렬합니다. 

<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) |  ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(log_n)](https://latex.codecogs.com/svg.image?O(\log&space;n)) |



<br />
<br />

## Cycle Sort

<br />
순환 정렬은 in-place 정렬이자 불안정 정렬 알고리즘으로, 배열에 대해 총 쓰기(write) 수의 관점에서 가장 적게 쓰기 위한 비교 정렬 알고리즘 입니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Odd-Even Sort

<br />
컴퓨팅 시스템 관점에서 홀짝 정렬은 원래 상호 로컬 연결의 병렬 프로세서에서 사용하기 위해 개발된 비교적 간단한 정렬 알고리즘입니다

<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2)](https://latex.codecogs.com/svg.image?O(n)) | Yes | Yes | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |



<br />
<br />

## Odd-Even Merge Sort

<br>

홀짝 병합 정렬은 사이즈는 
<span> ![O(n(log_n)^2)](https://latex.codecogs.com/svg.image?O(n(\log&space;n)^2)) </span>   , 깊이는   <span> ![O((log_n)^2)](https://latex.codecogs.com/svg.image?O((\log&space;n)^2)) </span>    를 갖는 네트워크 정렬을 위해 Ken Batcher 정렬 알고리즘 입니다.


<br />

### COMPLEXITY


| Type | Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| non-parallel | ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) | Yes | yes | total : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)), auxiliary : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) |
| parallel | ![O(log^2_n)](https://latex.codecogs.com/svg.image?O(\log^2&space;n)) | ![O(log^2_n)](https://latex.codecogs.com/svg.image?O(\log^2&space;n)) | ![O(log^2_n)](https://latex.codecogs.com/svg.image?O(\log^2&space;n)) | Yes | yes | total : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)), auxiliary : ![O(nlog^2_n)](https://latex.codecogs.com/svg.image?O(n\log^2&space;n)) |



<br />
<br />

## Comb Sort

<br />
콤 정렬은 원래 1980년 Włodzimierz Dobosiewicz와 Artur Borowy가 설계한 비교적 간단한 정렬 알고리즘으로, 나중에 Stephen Lacey와 Richard Box가 1991년에 재발견(이 때 "Combsort"라고 이름을 지정)했습니다. Comb 정렬은 셸 정렬과 유사한 방식으로 버블 정렬을 개선하여 성능을 향상 시킨 정렬 알고리즘입니다.
<br />

### COMPLEXITY


| Worst-Case | Average-Case | Best-Case | in-place | stable | Space Complexity |
| :-: | :-: | :-: | :-: | :-: | :-: |
| ![O(n^2)](https://latex.codecogs.com/svg.image?O(n^{2})) | ![O(n^2/p^2)](https://latex.codecogs.com/svg.image?O(n^{2}/2^{p})) | ![O(nlog_n)](https://latex.codecogs.com/svg.image?O(n\log&space;n)) | Yes | No | total : ![O(n)](https://latex.codecogs.com/svg.image?O(n)), auxiliary : ![O(1)](https://latex.codecogs.com/svg.image?O(1)) |
