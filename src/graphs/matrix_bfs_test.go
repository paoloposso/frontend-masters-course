package graphs

import "testing"

func Test_bfs(t *testing.T) {
	matrix := newMatrix(5)

	matrix[0][1] = 1
	matrix[0][2] = 1
	matrix[2][3] = 1
	matrix[3][4] = 1
	matrix[4][2] = 1

	out := bfs(matrix, 0, 3)

	if len(out) != 3 {
		t.Fail()
	}

	if out[0] != 3 {
		t.Fail()
	}
}
