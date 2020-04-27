package main

import (
	"lab1_sort_compare/internal"
)

type QuickSort struct {
}

func (q *QuickSort) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	q.sort(arr, 0, len(arr) - 1, swapper, comparer)
}

func (q *QuickSort) sort(arr []int, l, r int, swapper *internal.Swapper, comparer *internal.Comparer) {
	if l >= r {
		return
	}

	pivotIdx := l
	i := l + 1

	for j := l; j <= r; j++ {
		if comparer.Compare(arr[j], arr[pivotIdx]) {
			swapper.SwapArr(arr, i, j)
			i++
		}
	}

	swapper.SwapArr(arr, l, i-1)
	q.sort(arr, l, i-2, swapper, comparer)
	q.sort(arr, i, r, swapper, comparer)

}