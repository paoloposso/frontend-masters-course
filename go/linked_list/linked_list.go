package linkedlist

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	head   *Node[T]
	Length int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (linkedList *LinkedList[T]) InsertAt(value T, index int) error {
	if index > linkedList.Length {
		return errors.New("out of range")
	}
	newNode := &Node[T]{value: value, next: nil}

	next := linkedList.head
	var prev *Node[T]

	for i := 0; i < index; i++ {
		prev = next
		next = next.next
	}

	if prev != nil {
		prev.next = newNode
	}
	newNode.next = next
	linkedList.Length++
	return nil
}

func (linkedList *LinkedList[T]) RemoveAt(index int) error {
	if linkedList.Length < index+1 {
		return errors.New("out of range")
	}
	prev := linkedList.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	return nil
}

func (linkedList *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{value: value, next: nil}

	newNode.next = linkedList.head
	linkedList.head = newNode
	linkedList.Length++
}

func (linkedList *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{value: value, next: nil}

	if linkedList.head == nil {
		linkedList.head = newNode
		linkedList.Length++
		return
	}
	tail := linkedList.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = newNode
	linkedList.Length++
}

func (linkedList *LinkedList[T]) Print() {
	cur := linkedList.head
	for cur != nil {
		fmt.Printf("->%v", cur.value)
		cur = cur.next
	}
}

func (linkedList LinkedList[T]) Get(index int) (*T, error) {
	if index-1 > linkedList.Length {
		return nil, errors.New("out of range")
	}
	result := linkedList.head
	for i := 0; i < index; i++ {
		result = result.next
		if result == nil {
			return nil, errors.New("not found")
		}
	}
	return &result.value, nil
}
