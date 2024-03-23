package graphs

import "testing"

func TestCreateNodeAdjList(t *testing.T) {
	g := NewGraphAdjList()

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 2)
}
