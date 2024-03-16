package graphs

import (
	"errors"
	"fmt"
)

type Edge struct {
	Node   int
	Weight int
}

type Graph struct {
	AdjList map[int][]Edge
}

func NewGraph() Graph {
	return Graph{AdjList: map[int][]Edge{}}
}

func (g *Graph) CreateNode(node int) error {
	if g.NodeExists(node) {
		return errors.New("node already exists")
	}

	g.AdjList[node] = []Edge{}

	return nil
}

func (g *Graph) NodeExists(node int) bool {
	if _, ok := g.AdjList[node]; ok {
		return true
	}
	return false
}

func (g *Graph) EdgeExists(from, to int) bool {
	if edgesList, ok := g.AdjList[from]; ok {
		for _, edge := range edgesList {
			if edge.Node == to {
				return true
			}
		}
	}
	return false
}

func (g *Graph) CreateEdge(from, to int, weight int) error {
	if !g.NodeExists(from) {
		return fmt.Errorf("node %v does not exist", from)
	}
	if !g.NodeExists(to) {
		return fmt.Errorf("node %v does not exist", to)
	}
	if g.EdgeExists(from, to) {
		return fmt.Errorf("edge [%v, %v] already exists", from, to)
	}

	g.AdjList[from] = append(
		g.AdjList[from],
		Edge{Node: to, Weight: weight},
	)

	return nil
}
