package lca

import (
	"fmt"
	"log"
)

const maxn = 100010
const k = 26

func precalc(g map[int]map[int]bool, node, currheight int, e *[]int, h, f map[int]int) {
	if _, ok := h[node]; ok {
		return
	}
	h[node] = currheight
	if _, ok := f[node]; !ok {
		f[node] = len(*e)
	}
	*e = append(*e, node)
	for neigh := range g[node] {
		if _, ok := h[neigh]; ok {
			continue
		}
		precalc(g, neigh, currheight+1, e, h, f)
		*e = append(*e, node)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type nd struct {
	height, val int
}

func minnd(a, b *nd) *nd {
	if a == nil {
		if b != nil {
			log.Println("shu", b)
			return b
		}
		panic("err")
	}
	if b == nil {
		if a != nil {
			log.Println("shu", a)
			return a
		}
		panic("err")
	}
	if a.height < b.height {
		return a
	}
	return b
}

func Solve() {
	var n, m, a, b int
	g := map[int]map[int]bool{}
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		fmt.Scan(&b)
		if _, ok := g[b]; !ok {
			g[b] = map[int]bool{}
		}
		if _, ok := g[i]; !ok {
			g[i] = map[int]bool{}
		}
		g[i][b] = true
		g[b][i] = true
	}
	eulerpath := []int{}
	height := map[int]int{}
	first := map[int]int{}
	precalc(g, 0, 0, &eulerpath, height, first)

	logt := make([]int, maxn)
	st := make([][]*nd, maxn)
	logt[1] = 0
	for i := 2; i < maxn; i++ {
		logt[i] = logt[i/2] + 1
	}

	for i := 0; i < len(eulerpath); i++ {
		st[i] = make([]*nd, k+1)
		st[i][0] = &nd{height: height[eulerpath[i]], val: eulerpath[i]}
	}

	for j := 1; j <= k; j++ {
		for i := 0; i+(1<<uint(j)) <= len(eulerpath); i++ {
			st[i][j] = minnd(st[i][j-1], st[i+(1<<uint(j-1))][j-1])
		}
	}
	fmt.Scan(&m)
	fmt.Scan(&a)
	for i := 1; i < m; i++ {
		fmt.Scan(&b)
		R := max(first[a], first[b])
		L := min(first[a], first[b])
		j := logt[R-L+1]
		lca := minnd(st[L][j], st[R-(1<<uint(j))+1][j])
		a = lca.val
	}
	fmt.Println(a)
}
