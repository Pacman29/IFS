package main

import "lab1_sort_compare/internal"

type HeapSortFloyd struct {}

func (s *HeapSortFloyd) sink(arr []int, k, n int, swapper *internal.Swapper, comparer *internal.Comparer) {
	for 2*k <= n {
		j := 2*k
		if j < n && comparer.Compare(arr[j-1], arr[j]) {
			j++
		}
		if comparer.Compare( arr[k-1], arr[j-1]) {
			if j != k {
				swapper.SwapArr(arr, j-1, k-1)
			}
			k = j
		} else {
			break
		}
	}
}

func (s *HeapSortFloyd) floyd(arr []int, k, n int, swapper *internal.Swapper, comparer *internal.Comparer) {
	for 2*k <= n {
		j := 2*k
		if j < n && comparer.Compare(arr[j-1], arr[j]) {
			j++
		}
		if j != k {
			swapper.SwapArr(arr, j-1, k-1)
		}
		k = j
	}

	for k > 1 && comparer.Compare(arr[k/2 - 1], arr[k-1]) {
		swapper.SwapArr(arr, k-1, k/2-1)
		k = k/2
	}
}

func (s *HeapSortFloyd) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	n := len(arr)

	for i:= n/2; i >= 1; i-- {
		s.sink(arr, i, n, swapper, comparer)
	}

	for n > 1 {
		swapper.SwapArr(arr, 0, n-1)
		n--
		s.floyd(arr, 1, n, swapper, comparer)
	}
}
