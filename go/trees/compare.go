package trees

func compare(a, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Value != b.Value {
		return false
	}
	return compare(a.Left, b.Left) && compare(a.Right, b.Right)
}
