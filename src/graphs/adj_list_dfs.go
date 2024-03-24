package graphs

type AdjacencyList map[int][]edge

type edge struct {
	to, weight int
}

func NewGraphAdjList() AdjacencyList {
	return make(map[int][]edge)
}

func (gr AdjacencyList) addEdge(from, to int) {
	if _, ok := gr[from]; !ok {
		gr[from] = []edge{}
	}
	gr[from] = append(gr[from], edge{to: to, weight: 1})
}

func (gr AdjacencyList) dfs(source, needle int) (bool, []int) {
	seen := map[int]bool{}

	for i := range gr {
		seen[i] = false
	}

	path := &[]int{}

	itemFound := gr.dfsRecur(source, needle, seen, path)

	if itemFound {
		return true, append(*path, needle)
	}
	return false, []int{}
}

func (gr AdjacencyList) dfsRecur(curr, needle int, seen map[int]bool, path *[]int) bool {
	if curr == needle {
		return true
	}
	if seen[curr] {
		return false
	}

	seen[curr] = true

	*path = append(*path, curr)

	adj := gr[curr]

	for _, edge := range adj {
		if gr.dfsRecur(edge.to, needle, seen, path) {
			return true
		}
	}

	// pop
	*path = (*path)[:len(*path)-1]

	return false
}
