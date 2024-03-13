package queue

import "fmt"

type Queue[T any] struct {
	tail   *Item[T]
	head   *Item[T]
	Length int32
}

type Item[T any] struct {
	Data T
	next *Item[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{tail: nil, head: nil, Length: 0}
}

func (queue *Queue[T]) Queue(data T) {
	queue.Length++
	newItem := &Item[T]{
		Data: data,
		next: nil,
	}
	if queue.tail == nil {
		queue.tail = newItem
		queue.head = newItem
		return
	}
	queue.tail.next = newItem
	queue.tail = newItem
}

func (queue *Queue[T]) Dequeue() *Item[T] {
	result := queue.head
	if result == nil {
		return result
	}
	queue.head = queue.head.next
	queue.Length--
	return result
}

func (queue *Queue[T]) Peek() *Item[T] {
	return queue.head
}

func (queue Queue[T]) PrintQueue() {
	item := queue.head
	for item != nil {
		fmt.Printf("%v\n", item.Data)
		item = item.next
	}
}
