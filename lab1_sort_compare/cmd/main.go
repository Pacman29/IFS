package main

import (
	"fmt"
	"lab1_sort_compare/internal"
	"math/rand"
	"sort"
	"time"
)

func createShuffleArray(count int) []int {
	vals := make([]int, 0, count)
	for i := 0; i < count; i++ {
		vals = append(vals, i)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}

	return vals
}

func createShuffleRandomArray(count int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	vals := make([]int, 0, count)
	for i := 0; i < count; i++ {
		vals = append(vals, r.Intn(count))
	}

	for n := len(vals); n > 0; n-- {
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
	}

	return vals
}

func createSortArray(count int) []int {
	vals := make([]int, 0, count)
	for i := 0; i < count; i++ {
		vals = append(vals, i)
	}
	return vals
}

func createReversSortArray(count int) []int {
	vals := make([]int, 0, count)
	for i := count; i > 0; i-- {
		vals = append(vals, i)
	}
	return vals
}

func comparator(i, j int) bool {
	return i < j
}

func copyArr(arr []int) []int {
	arrCopy := append(arr[:0:0], arr...)
	return arrCopy
}

func printInfo(swapper *internal.Swapper, comparer *internal.Comparer) string {
	return fmt.Sprintf("swap count: %v; compare count: %v;", swapper.Count, comparer.Count)
}

func bubbleSortExec(arr []int) {
	bubble := BubbleSort{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	bubble.Sort(arr, swapper, comparer)

	fmt.Printf("bubble: %s\n", printInfo(swapper, comparer))
}

func insertionSortExec(arr []int) {
	insertion := InsertionSort{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	insertion.Sort(arr, swapper, comparer)

	fmt.Printf("insertion: %s\n", printInfo(swapper, comparer))
}

func insertionBinarySortExec(arr []int) {
	insertionBinary := InsertionBinarySort{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	insertionBinary.Sort(arr, swapper, comparer)

	fmt.Printf("insertion binary: %s \n", printInfo(swapper, comparer))
}

func heapSortExec(arr []int) {
	heap := HeapSort{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	heap.Sort(arr, swapper, comparer)

	fmt.Printf("heap: %s \n", printInfo(swapper, comparer))
}

func heapSortFloydExec(arr []int) {
	heap := HeapSortFloyd{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	heap.Sort(arr, swapper, comparer)

	fmt.Printf("heap floyd: %s\n", printInfo(swapper, comparer))
}

func shellSortExec(arr []int) {
	shell := ShellSort{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	shell.Sort(arr, swapper, comparer)

	fmt.Printf("shell: %s\n", printInfo(swapper, comparer))
}

func quickSortExec(arr []int) {
	quick := QuickSort{}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	quick.Sort(arr, swapper, comparer)

	fmt.Printf("quick: %s\n", printInfo(swapper, comparer))
}

func quickSortModExec(arr []int) {
	quick := QuickSortModified{ MinSize: 13, BigSize: 10_000, shell: &ShellSort{}}
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	quick.Sort(arr, swapper, comparer)

	fmt.Printf("quick mod: %s\n", printInfo(swapper, comparer))
}

type DataSort struct {
	arr []int
	swapper *internal.Swapper
	comparer *internal.Comparer
}

func (s *DataSort) Len() int {
	return len(s.arr)
}

func (s *DataSort) Less(i, j int) bool {
	return s.comparer.Compare(s.arr[i], s.arr[j])
}

func (s *DataSort) Swap(i, j int) {
	s.swapper.SwapArr(s.arr, i, j)
}

func quickSortDefaultExec(arr []int) {
	swapper := internal.NewSwapper()
	comparer := internal.NewComparer(comparator)

	sort.Sort(&DataSort{arr: arr, swapper: swapper, comparer: comparer})

	fmt.Printf("quick default: %s\n", printInfo(swapper, comparer))
}

func main() {
	arr := createReversSortArray(10_000)

	//fmt.Printf("init: arr: %v\n", arr)

	bubbleSortExec(copyArr(arr))
	insertionSortExec(copyArr(arr))
	insertionBinarySortExec(copyArr(arr))
	heapSortExec(copyArr(arr))
	heapSortFloydExec(copyArr(arr))
	shellSortExec(copyArr(arr))
	quickSortExec(copyArr(arr))
	quickSortModExec(copyArr(arr))
	quickSortDefaultExec(copyArr(arr))
}
