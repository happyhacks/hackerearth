package bigPAndParty

import (
	"container/list"
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)

func bfs(src int, g [][]int, d []int) {
	q := list.New()
	d[src] = 0
	for q.PushBack(src); q.Len() > 0; {
		next := q.Front()
		node := next.Value.(int)
		q.Remove(next)
		for _, neigh := range g[node] {
			if d[neigh] > d[node]+1 {
				d[neigh] = d[node] + 1
				q.PushBack(neigh)
			}
		}
	}
}

func Solve() {
	var n, m, a, b int
	fmt.Scan(&n)
	fmt.Scan(&m)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = MaxInt
	}
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	bfs(0, g, dist)
	for idx, l := range dist {
		if idx == 0 {
			continue
		}
		if l == MaxInt {
			fmt.Println(-1)
		} else {
			fmt.Println(l)
		}
	}
}
