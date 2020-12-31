package bigp

import (
	"container/heap"
	"container/list"
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

func bfs(g map[int]map[int]bool, dp []int, node int) {
	q := list.New()
	q.PushBack(node)
	dp[node] = 1
	for q.Len() > 0 {
		f := q.Front()
		if f == nil {
			break
		}
		nd := f.Value.(int)
		// log.Println(nd, dp)
		q.Remove(f)
		for neigh, vis := range g[nd] {
			if vis {
				dp[neigh] += dp[nd]
				q.PushBack(neigh)
				g[nd][neigh] = false
			}
		}
	}
}

func Solve() {
	var n, a, b int
	g := map[int]map[int]bool{}
	gh := map[int]*IntHeap{}
	indeg := map[int]int{}
	fmt.Scan(&n)
	dp := make([]int, n+1)
	for {
		fmt.Scan(&a)
		fmt.Scan(&b)
		if a == 0 && b == 0 {
			break
		}
		// if _, ok := g[b]; !ok {
		// 	g[b] = map[int]bool{}
		// }
		// g[b][a] = true

		if _, ok := g[a]; !ok {
			g[a] = map[int]bool{}
		}
		g[a][b] = true

		if _, ok := gh[a]; !ok {
			gh[a] = &IntHeap{}
			heap.Init(gh[a])
		}
		heap.Push(gh[a], b)
		indeg[b]++
	}
	// log.Println(g)
	visited := map[int]bool{}
	q := []int{}
	for i := n; i > 0; i-- {
		if _, ok := indeg[i]; !ok {
			dfs(i, gh, visited, &q)
		}
	}
	// log.Println(q)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		for neigh := range g[q[i]] {
			dp[q[i]] = dp[q[i]] + dp[neigh]
		}
	}
	// log.Println(dp)
	fmt.Println(dp[1])
}
