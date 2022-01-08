/*
  author : kdgyun

  link : https://st-lab.tistory.com
  link : https://github.com/kdgyun
*/

package main

type TimStack struct {
	array     []int
	runBase   []int
	runLen    []int
	stackSize int
}

func newTimStack(a []int) *TimStack {
	p := TimStack{}
	len := len(a)

	var stackLen int

	if len < 120 {
		stackLen = 5
	} else if len < 1542 {
		stackLen = 10
	} else if len < 119151 {
		stackLen = 19
	} else {
		stackLen = 40
	}
	p.runBase = make([]int, stackLen)
	p.runLen = make([]int, stackLen)
	p.array = a
	p.stackSize = 0

	return &p
}

func (p *TimStack) pushRun(base, len int) {
	p.runBase[p.stackSize] = base
	p.runLen[p.stackSize] = len
	p.stackSize++
}

func (p *TimStack) mergeForce() {
	for p.stackSize > 1 {
		if p.stackSize > 2 && p.runLen[p.stackSize-3] < p.runLen[p.stackSize-1] {
			p.mergeForTim(p.stackSize - 3)
		} else {
			p.mergeForTim(p.stackSize - 2)
		}
	}
}

func (p *TimStack) mergeBase() {
	for p.stackSize > 1 {
		if (p.stackSize > 2 && p.runLen[p.stackSize-3] <= p.runLen[p.stackSize-2]+p.runLen[p.stackSize-1]) || (p.stackSize > 3 && p.runLen[p.stackSize-4] <= p.runLen[p.stackSize-3]+p.runLen[p.stackSize-2]) {

			if p.runLen[p.stackSize-3] < p.runLen[p.stackSize-1] {
				p.mergeForTim(p.stackSize - 3)
			} else {
				p.mergeForTim(p.stackSize - 2)
			}

		} else if p.runLen[p.stackSize-2] <= p.runLen[p.stackSize-1] {
			p.mergeForTim(p.stackSize - 2)
		} else {
			break
		}
	}
}

func (p *TimStack) mergeForTim(idx int) {
	start1 := p.runBase[idx]
	length1 := p.runLen[idx]
	start2 := p.runBase[idx+1]
	length2 := p.runLen[idx+1]

	p.runLen[idx] = length1 + length2

	if idx == (p.stackSize - 3) {
		p.runBase[idx+1] = p.runBase[idx+2]
		p.runLen[idx+1] = p.runLen[idx+2]
	}
	p.stackSize--

	lo := p.gallopRight(p.array[start2], p.array, start1, length1)
	if length1 == lo {
		return
	}
	start1 += lo
	length1 -= lo

	hi := p.gallopLeft(p.array[start1+length1-1], p.array, start2, length2)
	if hi == 0 {
		return
	}

	length2 = hi
	if length1 <= length2 {
		p.mergeLo(start1, length1, start2, length2)
	} else {
		p.mergeHi(start1, length1, start2, length2)
	}
}

func (p *TimStack) mergeLo(start1, length1, start2, length2 int) {
	temp := make([]int, length1)
	copy(temp[:length1], p.array[start1:start1+length1])

	insertIdx := start1
	runBIdx := start2
	tempIdx := 0

	leftRemain := length1
	rightRemain := length2

	for leftRemain != 0 && rightRemain != 0 {
		if p.array[runBIdx] < temp[tempIdx] {
			p.array[insertIdx] = p.array[runBIdx]
			insertIdx++
			runBIdx++
			rightRemain--
		} else {
			p.array[insertIdx] = temp[tempIdx]
			insertIdx++
			tempIdx++
			leftRemain--
		}
	}
	if leftRemain != 0 {
		copy(p.array[insertIdx:], temp[tempIdx:tempIdx+leftRemain])
	} else {
		copy(p.array[insertIdx:], p.array[runBIdx:runBIdx+rightRemain])
	}
}

func (p *TimStack) mergeHi(start1, length1, start2, length2 int) {
	temp := make([]int, length2)
	copy(temp, p.array[start2:start2+length2])

	insertIdx := start2 + length2 - 1
	runAIdx := start1 + length1 - 1
	tempIdx := length2 - 1

	leftRemain := length1
	rightRemain := length2

	for leftRemain != 0 && rightRemain != 0 {

		if p.array[runAIdx] > temp[tempIdx] {
			p.array[insertIdx] = p.array[runAIdx]
			insertIdx--
			runAIdx--
			leftRemain--
		} else {
			p.array[insertIdx] = temp[tempIdx]
			insertIdx--
			tempIdx--
			rightRemain--
		}
	}
	if rightRemain != 0 {
		copy(p.array[start1:], temp[:rightRemain])
	}
}

func (p *TimStack) gallopRight(key int, array []int, base, lenOfA int) int {
	lo := 0
	hi := 1

	if key < array[base] {
		return 0
	}

	maxLen := lenOfA
	for hi < maxLen && array[base+hi] <= key {
		lo = hi
		hi = (hi << 1) + 1
		if hi <= 0 {
			hi = maxLen
		}
	}
	if hi > maxLen {
		hi = maxLen
	}

	lo++

	for lo < hi {
		mid := lo + ((hi - lo) >> 1)
		if key < array[base+mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

func (p *TimStack) gallopLeft(key int, array []int, base, lenOfB int) int {
	lo := 0
	hi := 1

	if key > array[base+lenOfB-1] {
		return lenOfB
	}

	startPointOfRun := base + lenOfB - 1
	maxLen := lenOfB

	for hi < maxLen && key <= array[startPointOfRun-hi] {
		lo = hi
		hi = (hi << 1) + 1

		if hi <= 0 {
			hi = maxLen
		}
	}

	if hi > maxLen {
		hi = maxLen
	}

	temp := lo
	lo = lenOfB - 1 - hi
	hi = lenOfB - 1 - temp

	lo++

	// binary search
	for lo < hi {
		mid := lo + ((hi - lo) >> 1)

		if key <= array[base+mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

func minRunLength(runSize int) int {
	r := 0
	for runSize >= 32 {
		r |= (runSize & 1)
		runSize >>= 1
	}
	return runSize + r
}

func Timsort(a []int) {
	timsort(a, 0, len(a))
}

func timsort(a []int, lo, hi int) {
	remain := hi - lo

	if remain < 2 {
		return
	}

	if remain < 32 {
		increaseRange := getAscending(a, lo, hi)
		binarySort(a, lo, hi, lo+increaseRange)
		return
	}

	ts := newTimStack(a)
	minRun := minRunLength(remain)

	for {
		runLen := getAscending(a, lo, hi)

		if runLen < minRun {
			var counts int
			if remain < minRun {
				counts = remain
			} else {
				counts = minRun
			}

			binarySort(a, lo, lo+counts, lo+runLen)
			runLen = counts
		}
		ts.pushRun(lo, runLen)
		ts.mergeBase()

		lo += runLen
		remain -= runLen
		if remain == 0 {
			break
		}
	}
	ts.mergeForce()
}
