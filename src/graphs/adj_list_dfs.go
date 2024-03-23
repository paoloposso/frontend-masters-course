package graphs

type AdjancencyList map[int][]int

func NewGraphAdjList() AdjancencyList {
	return map[int][]int{}
}

func (a AdjancencyList) addEdge(from, to int) {
	if _, ok := a[from]; !ok {
		a[from] = []int{}
	}
	a[from] = append(a[from], to)
}
