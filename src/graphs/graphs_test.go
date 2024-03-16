package graphs

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	g := NewGraph()

	g.CreateNode(1)
	g.CreateNode(2)

	a(g)

	nu := 1
	b(&nu)

	g.CreateEdge(1, 2, 1)
}

func a(g Graph) {
	g.CreateEdge(2, 1, 1)
	g.CreateNode(3)
	g.CreateEdge(3, 1, 1)
	g.CreateEdge(2, 3, 1)
	g.CreateEdge(3, 2, 1)
}

func b(b *int) {
	*b = 4
}
