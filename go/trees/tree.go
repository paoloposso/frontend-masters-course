package trees

type BinaryTree struct {
	Head *BinaryNode
}

type BinaryNode struct {
	Value int32
	Left  *BinaryNode
	Right *BinaryNode
}
