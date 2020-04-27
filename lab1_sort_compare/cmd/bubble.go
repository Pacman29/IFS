package main

import (
	"lab1_sort_compare/internal"
)

type BubbleSort struct {
}

func (b *BubbleSort) Sort(arr []int, swapper *internal.Swapper, comparer *internal.Comparer) {
	t := true
	n := len(arr)
	for i := 0; t && i < n - 1; i++ {
		t = false
		for j := 0; j < n - i - 1; j++ {
			if !comparer.Compare((arr)[j], (arr)[j+1]) {
				swapper.SwapArr(arr, j, j+1)
				t = true
			}
		}
	}
}


