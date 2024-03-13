package trees

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	head1 := createTree()
	head2 := createTree()

	res := compare(&head1, &head2)

	if res == false {
		t.Fail()
	}

	head2 = createTree2()

	res = compare(&head1, &head2)

	if res == true {
		t.Fail()
	}
}

func TestPreOrderSearch(t *testing.T) {

	head := createTree()

	fmt.Printf("%v", PreOrderSearch(head))
}

func TestBfs(t *testing.T) {

	head := createTree()

	result := Bfs(head, 20)

	if result == false {
		t.Fail()
	}

	result = Bfs(head, 800)

	if result == true {
		t.Fail()
	}

	result = Bfs(head, 18)

	if result == false {
		t.Fail()
	}
}

func TestDfs(t *testing.T) {

	head := createTree()

	result := dfs(&head, 20)

	if result == false {
		t.Fail()
	}

	result = Bfs(head, 800)

	if result == true {
		t.Fail()
	}

	result = Bfs(head, 18)

	if result == false {
		t.Fail()
	}
}

func createTree() BinaryNode {
	tr := BinaryTree{}

	tr.Head = &BinaryNode{Value: 1}
	tr.Head.Left = &BinaryNode{Value: 5}
	tr.Head.Right = &BinaryNode{Value: 20}

	curr := tr.Head.Left
	curr.Left = &BinaryNode{Value: 18}
	curr.Right = &BinaryNode{Value: 30}

	curr = curr.Right
	curr.Left = &BinaryNode{Value: 2}
	curr.Right = &BinaryNode{Value: 4}

	curr = curr.Right
	curr.Left = &BinaryNode{Value: 50}
	curr.Right = &BinaryNode{Value: 44}

	curr = tr.Head.Right
	curr.Left = &BinaryNode{Value: 18}
	curr.Right = &BinaryNode{Value: 30}

	curr = curr.Left
	curr.Left = &BinaryNode{Value: 1}
	curr.Right = &BinaryNode{Value: 11}

	curr = curr.Right
	curr.Left = &BinaryNode{Value: 1}
	curr.Right = &BinaryNode{Value: 11}

	curr = curr.Left
	curr.Left = &BinaryNode{Value: 18}
	curr.Right = &BinaryNode{Value: 55}

	return *tr.Head
}

func createTree2() BinaryNode {
	tr := BinaryTree{}

	tr.Head = &BinaryNode{Value: 1}
	tr.Head.Left = &BinaryNode{Value: 5}
	tr.Head.Right = &BinaryNode{Value: 20}

	curr := tr.Head.Left
	curr.Left = &BinaryNode{Value: 18}
	curr.Right = &BinaryNode{Value: 30}

	curr = curr.Right
	curr.Left = &BinaryNode{Value: 2}
	curr.Right = &BinaryNode{Value: 4}

	curr = curr.Right
	curr.Left = &BinaryNode{Value: 3}
	curr.Right = &BinaryNode{Value: 44}

	curr = tr.Head.Right
	curr.Left = &BinaryNode{Value: 18}
	curr.Right = &BinaryNode{Value: 30}

	curr = curr.Left
	curr.Left = &BinaryNode{Value: 1}
	curr.Right = &BinaryNode{Value: 11}

	curr = curr.Right
	curr.Left = &BinaryNode{Value: 1}
	curr.Right = &BinaryNode{Value: 11}

	curr = curr.Left
	curr.Left = &BinaryNode{Value: 18}
	curr.Right = &BinaryNode{Value: 55}

	return *tr.Head
}

func TestInsertMultiple(t *testing.T) {
	var head *BinaryNode = &BinaryNode{Value: 5}

	insert(head, 3)
	insert(head, 8)
	insert(head, 4)
	insert(head, 1)

	preOrderArr := PreOrderSearch(*head)

	fmt.Printf("%v", preOrderArr)
}
