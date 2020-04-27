package main

import (
	"lab1_sort_compare/internal"
	"testing"
)


func createShuffleRandomArrays(n, count int) [][]int {
	var vals = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		vals = append(vals, createShuffleRandomArray(count))
	}
	return vals
}

func Benchmark_Sort_10(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createSortArray(10)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_Sort_20(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createSortArray(20)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_Sort_40(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createSortArray(40)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_Sort_100(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createSortArray(100)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_Sort_5000(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createSortArray(5000)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_Sort_10000(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createSortArray(10000)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

//// Revers

func Benchmark_ReversSort_10(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createReversSortArray(10)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_ReversSort_20(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createReversSortArray(20)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_ReversSort_40(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createReversSortArray(40)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_ReversSort_100(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createReversSortArray(100)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_ReversSort_5000(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createReversSortArray(5000)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_ReversSort_10000(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createReversSortArray(10000)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arr := copyArr(benchmarkData)
				b.StartTimer()
				bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
				b.StopTimer()
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

// Random

func Benchmark_RandomSort_10(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createShuffleRandomArrays(10, 10)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bench := range benchmarkData {
					arr := copyArr(bench)
					b.StartTimer()
					bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
					b.StopTimer()
				}
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_RandomSort_20(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createShuffleRandomArrays(10, 20)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bench := range benchmarkData {
					arr := copyArr(bench)
					b.StartTimer()
					bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
					b.StopTimer()
				}
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_RandomSort_40(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createShuffleRandomArrays(10, 40)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bench := range benchmarkData {
					arr := copyArr(bench)
					b.StartTimer()
					bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
					b.StopTimer()
				}
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_RandomSort_100(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createShuffleRandomArrays(10, 100)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bench := range benchmarkData {
					arr := copyArr(bench)
					b.StartTimer()
					bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
					b.StopTimer()
				}
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_RandomSort_5000(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createShuffleRandomArrays(10, 5000)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bench := range benchmarkData {
					arr := copyArr(bench)
					b.StartTimer()
					bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
					b.StopTimer()
				}
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}

func Benchmark_RandomSort_10000(b *testing.B) {
	b.StopTimer()
	benchmarkSorters := []struct {
		name string
		sort internal.Sorter
		comparer *internal.Comparer
		swapper  *internal.Swapper
	}{
		{"BubbleSort",&BubbleSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionSort",&InsertionSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"InsertionBinarySort",&InsertionBinarySort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSort",&HeapSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"HeapSortFloyd",&HeapSortFloyd{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"ShellSort",&ShellSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSort",&QuickSort{}, internal.NewComparer(comparator), internal.NewSwapper()},
		{"QuickSortModified",&QuickSortModified{MinSize: 13, BigSize: 2500, shell: &ShellSort{}}, internal.NewComparer(comparator), internal.NewSwapper()},
	}

	benchmarkData := createShuffleRandomArrays(10, 10_000)

	b.ResetTimer()
	for _, bmSorter := range benchmarkSorters {
		b.Run(bmSorter.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, bench := range benchmarkData {
					arr := copyArr(bench)
					b.StartTimer()
					bmSorter.sort.Sort(arr, bmSorter.swapper, bmSorter.comparer)
					b.StopTimer()
				}
			}
		})
	}

	for _, sorter := range benchmarkSorters {
		sorter.swapper.Count /= 100
		sorter.comparer.Count /= 100
		b.Logf("%v: %v", sorter.name, printInfo(sorter.swapper, sorter.comparer))
	}
}