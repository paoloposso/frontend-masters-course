package stack

type Stack[T any] struct {
	Length int32
	Top    *Item[T]
}

type Item[T any] struct {
	Value T
	Next  *Item[T]
}

func (stack *Stack[T]) Push(value T) {
	stack.Length++
	newItem := Item[T]{Value: value, Next: stack.Top}
	stack.Top = &newItem
}

func (stack *Stack[T]) Pop() *Item[T] {
	if stack.Top == nil {
		return nil
	}
	stack.Length--
	result := stack.Top
	stack.Top = stack.Top.Next
	return result
}

func (stack *Stack[T]) Peak() *Item[T] {
	return stack.Top
}
