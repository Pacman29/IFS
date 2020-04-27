package main

import "lab1_sort_compare/internal"

type HeapSort struct {}

func (s *HeapSort) heapify(arr []int, n, i int, swapper *internal.Swapper, comparer *internal.Comparer) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && comparer.Compare(arr[largest], arr[l]) {
		largest = l
	}

	if r < n && comparer.Compare(arr[largest], arr[r]) {
		largest = r
	}

	if largest != i {
		swapper.SwapArr(arr, i, largest)
		s.heapify(arr, n, largest, swapper, comparer)
	}
}

func (s *HeapSort) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	n := len(arr)

	for i:= n/2 -1; i >= 0; i-- {
		s.heapify(arr, n, i, swapper, comparer)
	}

	for i:= n-1; i >= 0; i-- {
		swapper.SwapArr(arr, 0, i)
		s.heapify(arr, i, 0, swapper, comparer)
	}
}
