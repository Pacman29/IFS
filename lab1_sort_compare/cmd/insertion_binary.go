package main

import (
	"lab1_sort_compare/internal"
)

type InsertionBinarySort struct {}

func (i *InsertionBinarySort) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	n := len(arr)
	for i := 1; i < n; i++ {
		l := 0
		r := i

		for l < r {
			m := l + (r - l)/2
			if comparer.Compare(arr[i], arr[m]) {
				r = m
			} else {
				l = m + 1
			}
		}

		for j := i; j > l; j-- {
			swapper.SwapArr(arr, j , j-1)
		}
	}
}

