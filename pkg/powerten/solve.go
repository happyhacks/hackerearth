package powerten

import (
	"fmt"
)

const maxn = 1000001
const k = 25

func log2(n int) int {
	s := 1
	for i := 0; i < 32; i++ {
		if n%s != 0 {
			return i - 1
		}
		s *= 2
	}
	return -1
}

func log5(n int) int {
	s := 1
	for i := 0; i < 32; i++ {
		if n%s != 0 {
			return i - 1
		}
		s *= 5
	}
	return -1
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
			// log.Println(b)
			return b
		}
		panic("err")
	}
	if b == nil {
		if a != nil {
			// log.Println(a)
			return a
		}
		panic("err")
	}
	if a.height < b.height {
		return a
	}
	return b
}

func precalc(g map[int]map[int]int, node, currheight, cl2, cl5 int, e *[]int, h, f, l2, l5 map[int]int) {
	if _, ok := h[node]; ok {
		return
	}
	h[node] = currheight
	l2[node] = cl2
	l5[node] = cl5
	if _, ok := f[node]; !ok {
		f[node] = len(*e)
	}
	*e = append(*e, node)
	for neigh, wt := range g[node] {
		if _, ok := h[neigh]; ok {
			continue
		}
		precalc(g, neigh, currheight+1, cl2+log2(wt), cl5+log5(wt), e, h, f, l2, l5)
		*e = append(*e, node)
	}
}

func Solve() {
	var n, m, a, b, w int
	g := map[int]map[int]int{}
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&w)
		if _, ok := g[a]; !ok {
			g[a] = map[int]int{}
		}
		if _, ok := g[b]; !ok {
			g[b] = map[int]int{}
		}
		g[a][b] = w
		g[b][a] = w
	}
	eulerpath := []int{}
	height := map[int]int{}
	first := map[int]int{}
	l2 := map[int]int{}
	l5 := map[int]int{}
	precalc(g, 1, 0, 0, 0, &eulerpath, height, first, l2, l5)
	// log.Println(first)
	// log.Println(eulerpath)
	// log.Println(height)

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
		for i := 0; i+(1<<j) <= len(eulerpath); i++ {
			st[i][j] = minnd(st[i][j-1], st[i+(1<<(j-1))][j-1])
		}
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		// if a == 1 {
		// 	log.Println(l2[b], l5[b])
		// 	fmt.Println(min(l2[b], l5[b]))
		// 	continue
		// }
		// if b == 1 {
		// 	log.Println(l2[a], l5[a])
		// 	fmt.Println(min(l2[a], l5[a]))
		// 	continue
		// }

		R := max(first[a], first[b])
		L := min(first[a], first[b])
		j := logt[R-L+1]
		lca := minnd(st[L][j], st[R-(1<<j)+1][j])
		// log.Println(lca)
		fmt.Println(min(l2[a]+l2[b]-l2[lca.val]-l2[lca.val], l5[a]+l5[b]-l5[lca.val]-l5[lca.val]))
	}

}
