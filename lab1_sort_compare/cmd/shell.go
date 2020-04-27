package main

import (
	"lab1_sort_compare/internal"
)

type ShellSort struct {
}

func (b *ShellSort) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	b.sort(arr, len(arr), swapper, comparer)
}

func (b *ShellSort) sort(arr []int, n int, swapper *internal.Swapper, comparer *internal.Comparer) {
	for i := n/2; i>0; i /= 2 {
		for j := i; j < n; j++ {
			for k := j; k >= i && comparer.Compare(arr[k], arr[k-1]); k -= i {
				swapper.SwapArr(arr, k, k - i)
			}
		}
	}
}
