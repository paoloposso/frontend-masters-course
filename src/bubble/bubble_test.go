package bubble

import (
	"testing"
)

func TestSortRecur(t *testing.T) {
	arr := []int{8, 9, 10, 5, 3, 6}
	SortRecur(arr, len(arr))

	if arr[3] < arr[2] {
		t.Fail()
	}
}

func TestSort(t *testing.T) {
	arr := []int{8, 9, 10, 5, 3, 6}
	SortRecur(arr, len(arr))

	if arr[3] < arr[2] {
		t.Fail()
	}
}
