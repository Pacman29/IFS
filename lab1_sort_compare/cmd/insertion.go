package main

import "lab1_sort_compare/internal"

type InsertionSort struct {}

func (i *InsertionSort) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j:= i-1; j >= 0 && !comparer.Compare(arr[j], arr[j+1]); j-- {
			swapper.SwapArr(arr, j, j+1)
		}
	}
}
