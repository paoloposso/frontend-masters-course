package graphs

import (
	"testing"
)

func TestCreateNodeAdjList(t *testing.T) {
	g := NewGraphAdjList()

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	g.addEdge(4, 2)

	if len(g[1]) != 2 {
		t.Fail()
	}
}

func TestDfs(t *testing.T) {
	g := NewGraphAdjList()

	g.addEdge(1, 2)
	g.addEdge(2, 4)
	g.addEdge(2, 5)
	g.addEdge(4, 5)
	g.addEdge(4, 6)
	g.addEdge(6, 3)

	found, path := g.dfs(1, 3)

	if found == false {
		t.Fail()
	}
	if path[len(path)-1] != 3 {
		t.Fail()
	}
	if path[0] != 1 {
		t.Fail()
	}
}

func TestDfsEmpty(t *testing.T) {
	g := NewGraphAdjList()

	found, path := g.dfs(1, 3)

	if found == true {
		t.Fail()
	}
	if len(path) > 0 {
		t.Fail()
	}
}
