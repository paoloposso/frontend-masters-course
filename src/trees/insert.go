package trees

func insert(head *BinaryNode, value int32) {
	newNode := &BinaryNode{Value: value}
	if head == nil {
		head = newNode
		return
	}

	insertRecur(head, newNode)
}

func insertRecur(curr, newNode *BinaryNode) {
	if newNode.Value <= curr.Value {
		if curr.Left == nil {
			curr.Left = newNode
			return
		}
		insertRecur(curr.Left, newNode)
	} else if newNode.Value >= curr.Value {
		if curr.Right == nil {
			curr.Right = newNode
			return
		}
		insertRecur(curr.Right, newNode)
	}
}
