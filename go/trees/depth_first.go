package trees

func dfs(head *BinaryNode, needle int32) bool {
	return search(head, needle)
}

func search(current *BinaryNode, needle int32) bool {
	if current == nil {
		return false
	}
	if current.Value == needle {
		return true
	}

	if current.Value > needle {
		return search(current.Left, needle)
	} else {
		return search(current.Right, needle)
	}
}
