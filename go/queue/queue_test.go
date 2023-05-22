package queue

import (
	"testing"
)

func TestPush(t *testing.T) {
	q := NewQueue[string]()
	q.Queue("aaaa")
	q.Queue("bbbb")
	q.Queue("cccc")
	q.Queue("dddd")
	q.Queue("eeee")

	v := q.Dequeue()
	if v.Data != "aaaa" {
		t.Fail()
	}

	v = q.Dequeue()
	if v.Data != "bbbb" {
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	q := NewQueue[string]()
	q.Queue("aaaa")
	q.Queue("bbbb")
	q.Queue("cccc")
	q.Queue("dddd")
	q.Queue("eeee")

	v := q.Dequeue()
	if v.Data != "aaaa" {
		t.Fail()
	}

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	if q.Length != 0 {
		t.Fail()
	}
}
