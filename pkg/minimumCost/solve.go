package minimumCost

import (
	"container/list"
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func bfs(node int, g map[int]map[int]int, d []int) {
	q := list.New()
	d[node] = 0
	for q.PushBack(node); q.Len() > 0; {
		next := q.Front()
		neigh := next.Value.(int)
		q.Remove(next)
		for v, dist := range g[neigh] {
			if d[v] > d[neigh]+dist {
				d[v] = d[neigh] + dist
				if dist == 0 {
					q.PushFront(v)
				} else {
					q.PushBack(v)
				}
			}
		}
	}
}

func solve() {
	var n, a int
	fmt.Scan(&n)
	p := make([]int, n+1)
	g := map[int]map[int]int{}
	for i := 1; i <= n; i++ {
		g[i] = map[int]int{}
		if i != n {
			g[i][i+1] = 1
		}
		if i != 1 {
			g[i][i-1] = 1
		}
	}
	d := make([]int, n+1, n+1)
	for i := 1; i <= n; i++ {
		d[i] = MaxInt
		fmt.Scan(&a)
		p[i] = a
		if i == a {
			continue
		}
		g[i][a] = 0
	}
	// log.Println(g)
	bfs(1, g, d)
	fmt.Println(d[n])
}

func Solve() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
