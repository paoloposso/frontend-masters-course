package graphs

type AdjacencyList map[int][]edge

type edge struct {
	to, weight int
}

const infinity int = 999999999999999999

func NewGraphAdjList() AdjacencyList {
	return make(map[int][]edge)
}

func (gr AdjacencyList) AddEdge(from, to, weight int) {
	if _, ok := gr[from]; !ok {
		gr[from] = []edge{}
	}
	gr[from] = append(gr[from], edge{to: to, weight: weight})

	if _, ok := gr[to]; !ok {
		for c := len(gr) - 1; c <= to; c++ {
			if _, ok := gr[c]; !ok {
				gr[c] = []edge{}
			}
		}
	}
}

func (gr AdjacencyList) Dfs(source, needle int) (bool, []int) {
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

func (graph AdjacencyList) Dijkstra(source, sink int) []int {
	seen := make([]bool, len(graph))

	for i := range seen {
		seen[i] = false
	}

	dists := make([]int, len(graph))

	for i := range dists {
		dists[i] = infinity
	}

	prev := make([]int, len(graph))

	for i := range prev {
		prev[i] = -1
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := graph[curr]

		for _, edge := range adjs {
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + edge.weight
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	curr := sink
	path := []int{}

	for prev[curr] != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	result := []int{}

	for i := range path {
		result = append(result, path[len(path)-1-i])
	}

	return append([]int{source}, result...)
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, s := range seen {
		if !s && dists[i] < infinity {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	result := -1
	lowestDist := infinity

	for i, s := range seen {
		if s {
			continue
		}

		if lowestDist > dists[i] {
			lowestDist = dists[i]
			result = i
		}
	}

	return result
}
