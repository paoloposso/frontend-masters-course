package graphs

type AdjacencyMatrix [][]int

func newMatrix(size int) AdjacencyMatrix {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return matrix
}

func bfs(graph AdjacencyMatrix, source, needle int) []int {
	graphLen := len(graph)
	seen := make([]bool, graphLen)
	for i := range seen {
		seen[i] = false
	}
	prev := make([]int, graphLen)
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true

	q := []int{source}

	for {
		curr := q[0]
		q = q[1:]

		if curr == needle {
			break
		}

		adj := graph[curr]

		for i, weight := range adj {
			if seen[i] {
				continue
			}
			if weight == 0 {
				continue
			}
			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}

		if len(q) == 0 {
			break
		}
	}

	curr := needle
	out := []int{}

	for {
		out = append(out, curr)
		curr = prev[curr]

		if prev[curr] == -1 {
			break
		}
	}

	return append(out, source)
}
