package trees

func Bfs(head BinaryNode, needle int32) bool {
	q := []BinaryNode{head}

	for len(q) > 0 {
		// pop
		curr := q[0]
		q = q[1:]

		if curr.Value == needle {
			return true
		}

		//push children
		if curr.Left != nil {
			q = append(q, *curr.Left)
		}
		if curr.Right != nil {
			q = append(q, *curr.Right)
		}
	}

	return false
}
