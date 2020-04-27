package internal

type Sorter interface{
	Sort(arr []int, swapper *Swapper, comparer *Comparer)
}