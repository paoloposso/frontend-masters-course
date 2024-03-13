package stack

import "testing"

func TestPush(t *testing.T) {
	q := Stack[int32]{}
	q.Push(3)
	q.Push(2)
	q.Push(1)
	q.Push(4)
	v := q.Pop()
	if v.Value != 4 {
		t.Fail()
	}
	if v = q.Peak(); v.Value != 1 {
		t.Fail()
	}
	q.Push(5)
	if v = q.Pop(); v.Value != 5 {
		t.Fail()
	}
	if q.Length != 3 {
		t.Fail()
	}
}
