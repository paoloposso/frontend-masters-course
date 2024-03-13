package heap

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	h := NewHeap()

	h.Add(50)
	h.Add(71)
	h.Add(100)
	h.Add(101)
	h.Add(80)
	h.Add(200)
	h.Add(101)

	if lastItem, err := h.GetLastItem(); err != nil {
		t.Error("Error:", err)
	} else if *lastItem != 101 {
		t.Error("Last item is not correct")
	}
}

func TestInsertionOfSmallestValue(t *testing.T) {
	h := NewHeap()

	h.Add(50)
	h.Add(71)
	h.Add(100)
	h.Add(101)
	h.Add(80)
	h.Add(200)
	h.Add(101)
	h.Add(3)

	if h.arr[0] != 3 {
		t.Error("First item should be 3")
	}
}

func TestRemoveMinimum(t *testing.T) {
	h := NewHeap()

	h.Add(50)
	h.Add(71)
	h.Add(100)
	h.Add(101)
	h.Add(80)
	h.Add(200)
	h.Add(101)
	h.Add(3)

	val, _ := h.RemoveMinimumItem()

	if *val != 3 {
		t.Error("Value should be 3")
	}

	if h.arr[0] != 50 {
		t.Error("Minimum value should be 50")
	}

	val, _ = h.RemoveMinimumItem()

	if *val != 50 {
		t.Error("Value should be 50")
	}

	if h.arr[0] != 71 {
		t.Error("Minimum value should be 71")
	}
}
