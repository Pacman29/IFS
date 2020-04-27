package main

import (
	"lab1_sort_compare/internal"
	"math/rand"
)

type QuickSortModified struct {
	MinSize int
	BigSize int
	shell   *ShellSort
}

func (q *QuickSortModified) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	q.sort(arr, 0, len(arr), swapper, comparer)
}

type pair struct {
	data int
	pos  int
}

func (q *QuickSortModified) compare(i, j pair, comparer *internal.Comparer) bool {
	return comparer.Compare(i.data, j.data)
}

func (q *QuickSortModified) swapArr(arr []pair, i, j int, swapper *internal.Swapper) {
	swapper.Count++
	arr[i], arr[j] = arr[j], arr[i]
}

func (q *QuickSortModified) pairSort(arr []pair, n int, swapper *internal.Swapper, comparer *internal.Comparer) {
	for i := n/2; i>0; i /= 2 {
		for j := i; j < n; j++ {
			for k := j; k >= i && q.compare(arr[k], arr[k-1], comparer); k -= i {
				q.swapArr(arr, k, k - i, swapper)
			}
		}
	}
}

func (q *QuickSortModified) choosePivot(arr []int, n int, swapper *internal.Swapper, comparer *internal.Comparer) int {
	if n > q.BigSize {
		pivotCount := 10
		pivotArr := make([]pair, 0, pivotCount)
		for i := 0; i<pivotCount; i++ {
			idx := i * n/pivotCount
			pivotArr = append(pivotArr, pair{data: arr[idx], pos: idx })
		}
		q.pairSort(pivotArr, pivotCount, swapper, comparer)
		return pivotArr[pivotCount /2].pos
	}

	middle := n/2
	if comparer.Compare(arr[0], arr[middle]) {
		if comparer.Compare(arr[middle], arr[n-1]) {
			return middle
		}
		if comparer.Compare(arr[n-1], arr[0]) {
			return 0
		}
		return n-1
	}

	if comparer.Compare(arr[n-1], arr[middle]) {
		return middle
	}

	if comparer.Compare(arr[0], arr[n-1]) {
		return n-1
	}
	return 0
}

func (q *QuickSortModified) choosePivot2(arr []int, n int, swapper *internal.Swapper, comparer *internal.Comparer) int {
	if n > q.BigSize {
		pivotCount := 10
		pivotArr := make([]pair, 0, pivotCount)
		for i := 0; i<pivotCount; i++ {
			idx := i * n/pivotCount
			pivotArr = append(pivotArr, pair{data: arr[idx], pos: idx })
		}
		q.pairSort(pivotArr, pivotCount, swapper, comparer)
		return pivotArr[pivotCount/2].pos
	}

	return rand.Intn(n)
}

func (q *QuickSortModified) medianOfThree(arr []int, m1, m0, m2 int, swapper *internal.Swapper, comparer *internal.Comparer) {
	// sort 3 elements
	if comparer.Compare(arr[m1], arr[m0]) {
		swapper.SwapArr(arr, m1, m0)
	}
	// data[m0] <= data[m1]
	if comparer.Compare(arr[m2], arr[m1]) {
		swapper.SwapArr(arr, m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if comparer.Compare(arr[m1], arr[m0]) {
			swapper.SwapArr(arr, m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func (q *QuickSortModified) choosePivot3(arr []int, n int, swapper *internal.Swapper, comparer *internal.Comparer) int {
	m := n >> 1
	if n > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := n / 8
		q.medianOfThree(arr, 0, s, 2*s, swapper, comparer)
		q.medianOfThree(arr, m, m-s, m+s, swapper, comparer)
		q.medianOfThree(arr, n-1, n-1-s, n-1-2*s, swapper, comparer)
	}
	q.medianOfThree(arr, 0, m, n-1, swapper, comparer)

	return 0
}

func (q *QuickSortModified) partition(arr []int, n int, swapper *internal.Swapper, comparer *internal.Comparer) int {
	middle := q.choosePivot(arr, n, swapper, comparer)
	if 0 != middle {
		swapper.SwapArr(arr, 0, middle)
	}
	pivot := 0

	i := n-1
	for j := i; j > pivot; j-- {
		if comparer.Compare(arr[pivot], arr[j]) {
			swapper.SwapArr(arr, i, j)
			i--
		}
	}
	if pivot != i {
		swapper.SwapArr(arr, pivot, i)
	}

	return i
}

func (q *QuickSortModified) sort(arr []int, begin, end int, swapper *internal.Swapper, comparer *internal.Comparer) {
	stack := make([]int, 0, end)
	stack = append(stack, begin, end)

	for len(stack) > 0 {
		begin, end, stack = stack[len(stack)-2], stack[len(stack)-1], stack[:len(stack)-2]
		if (end - begin) < q.MinSize {
			q.shell.sort(arr[begin:end], end - begin, swapper, comparer)
			continue
		}
		middle := q.partition(arr[begin:end], end - begin, swapper, comparer) + begin
		it, jt := middle, middle
		for it > begin +1 && comparer.Equal(arr[it], arr[it-1]) {
			it--
		}
		for jt < end -1 && comparer.Equal(arr[jt], arr[jt+1]) {
			jt++
		}

		stack = append(stack, begin, it, jt+1, end)
	}
}


