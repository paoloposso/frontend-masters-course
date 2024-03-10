package heap

import "errors"

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

		// if current is bigger than prev, bubble up is finished
		if h.arr[i] > h.arr[prevIndex] {
			return
		}

		// otherwise, swap
		h.arr[i], h.arr[prevIndex] = h.arr[prevIndex], h.arr[i]
		i = prevIndex
	}
}

func (h *Heap) Remove(val int32) (*int32, error) {
	if len(h.arr) == 0 {
		return nil, errors.New("heap is empty")
	}
	removed := h.arr[0]

	// heap down

	return &removed, nil
}

func (h *Heap) GetLastItem() (*int32, error) {
	if len(h.arr) == 0 {
		return nil, errors.New("heap is empty")
	}
	return &h.arr[len(h.arr)-1], nil
}
