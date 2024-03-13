package heap

import (
	"errors"
)

type Heap struct {
	arr []int32
}

func NewHeap[h Heap]() *Heap {
	return &Heap{arr: []int32{}}
}

func (h *Heap) Add(val int32) {
	h.arr = append(h.arr, val)

	prevIndex := -1
	i := len(h.arr) - 1

	for i > 0 {
		if i%2 == 0 {
			prevIndex = (i / 2) - 1
		} else {
			prevIndex = i / 2
		}

		if h.arr[i] > h.arr[prevIndex] {
			return
		}

		h.arr[i], h.arr[prevIndex] = h.arr[prevIndex], h.arr[i]
		i = prevIndex
	}
}

func (h *Heap) RemoveMinimumItem() (*int32, error) {
	if len(h.arr) == 0 {
		return nil, errors.New("heap is empty")
	}
	removed := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]

	h.arr = h.arr[:len(h.arr)-1]

	heapifyDownRecur(h.arr, 0, 1, 2)

	return &removed, nil
}

func heapifyDownRecur(arr []int32, curr, ch1, ch2 int) {
	if ch1 >= len(arr) {
		return
	}

	swapIndex := curr

	if arr[ch1] < arr[curr] {
		swapIndex = ch1
	}
	if ch2 < len(arr) && arr[ch2] < arr[swapIndex] {
		swapIndex = ch2
	}

	arr[curr], arr[swapIndex] = arr[swapIndex], arr[curr]

	heapifyDownRecur(arr, swapIndex, swapIndex*2+1, swapIndex*2+2)
}

func (h *Heap) GetLastItem() (*int32, error) {
	if len(h.arr) == 0 {
		return nil, errors.New("heap is empty")
	}
	return &h.arr[len(h.arr)-1], nil
}
