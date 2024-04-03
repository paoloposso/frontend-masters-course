package graphs

import (
	"fmt"
	"testing"
)

func TestCreateNodeAdjList(t *testing.T) {
	g := NewGraphAdjList()

	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 2, 1)

	if len(g[1]) != 2 {
		t.Fail()
	}
}

func TestDfs(t *testing.T) {
	g := NewGraphAdjList()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(4, 6, 1)
	g.AddEdge(6, 3, 1)

	found, path := g.Dfs(1, 3)

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

	found, path := g.Dfs(1, 3)

	if found == true {
		t.Fail()
	}
	if len(path) > 0 {
		t.Fail()
	}
}

func TestNonExistingValue(t *testing.T) {
	g := NewGraphAdjList()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 5, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(4, 6, 1)
	g.AddEdge(6, 3, 1)

	found, path := g.Dfs(1, 7)

	if found == true {
		t.FailNow()
	}
	if len(path) > 0 {
		t.FailNow()
	}
}

func TestDijkstra(t *testing.T) {
	g := NewGraphAdjList()

	g.AddEdge(0, 1, 8)
	g.AddEdge(0, 2, 2)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 4, 12)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 2)

	result := g.Dijkstra(0, 4)

	if fmt.Sprint(result) != "[0 2 3 4]" {
		t.Fail()
	}
}
