package main

import "fmt"

func quickSort(arr []int) {
	quickSortRecur(arr, 0, len(arr)-1)
}

func quickSortRecur(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	p := partition(arr, lo, hi)

	quickSortRecur(arr, lo, p-1)
	quickSortRecur(arr, p+1, hi)
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]

	ix := start - 1

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			ix++
			t := arr[ix]
			arr[ix] = arr[j]
			arr[j] = t
		}
	}

	t := arr[ix+1]
	arr[ix+1] = arr[end]
	arr[end] = t

	return ix + 1
}

func main() {
	arr := []int{4, 5, 8, 10, 16, 1, 7}

	quickSort(arr)

	fmt.Println(arr)
}
