package topo

import (
	"container/heap"
	"fmt"
)

// Slice of ints
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}

func dfs(src int, g map[int]*IntHeap, v map[int]bool, q *[]int) {
	// log.Println(src)
	if _, ok := v[src]; ok {
		return
	}

	v[src] = true
	if _, ok := g[src]; ok {
		ln := g[src].Len()
		for i := 0; i < ln; i++ {
			neigh := heap.Pop(g[src]).(int)
			// log.Println(src, neigh)
			if _, ok := v[neigh]; ok {
				continue
			}
			dfs(neigh, g, v, q)
		}
	}

	*q = append([]int{src}, *q...)
}

func Solve() {
	var n, m, a, b int
	g := map[int]*IntHeap{}
	indeg := map[int]int{}
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		if _, ok := g[a]; !ok {
			g[a] = &IntHeap{}
			heap.Init(g[a])
		}
		heap.Push(g[a], b)
		indeg[b]++
	}
	visited := map[int]bool{}
	q := []int{}
	for i := n; i > 0; i-- {
		if _, ok := indeg[i]; !ok {
			dfs(i, g, visited, &q)
			// break
		}
	}
	for _, i := range q {
		fmt.Printf("%d ", i)
	}
}
