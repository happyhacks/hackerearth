package fixtheroads

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(node, parent, timer int, g, gb map[int]map[int]bool, v map[int]bool, tin, low []int) int {
	v[node] = true
	t := timer + 1
	tin[node] = t
	low[node] = t
	for neigh := range g[node] {
		if neigh == parent {
			continue
		}
		if _, ok := v[neigh]; ok {
			low[node] = min(low[node], tin[neigh])
		} else {
			t = dfs(neigh, node, t, g, gb, v, tin, low)
			low[node] = min(low[node], low[neigh])
			if low[neigh] > tin[node] {
				if _, ok := gb[neigh]; !ok {
					gb[neigh] = map[int]bool{}
				}
				gb[neigh][node] = true

				if _, ok := gb[node]; !ok {
					gb[node] = map[int]bool{}
				}
				gb[node][neigh] = true
			}
		}
	}
	return t
}

func diadfs(node, parent, depth int, maxdepth, maxnode *int, g map[int]map[int]bool, v map[int]bool) {
	v[node] = true
	for neigh := range g[node] {
		if neigh == parent {
			continue
		}
		diadfs(neigh, node, depth+1, maxdepth, maxnode, g, v)
	}
	if depth > *maxdepth {
		*maxnode = node
		*maxdepth = depth
	}
}

func dfscc(node, cc int, g, gb map[int]map[int]bool, ccm map[int]int) {
	if _, ok := ccm[node]; ok {
		return
	}
	ccm[node] = cc
	for neigh := range g[node] {
		if _, ok := gb[node][neigh]; !ok {
			dfscc(neigh, cc, g, gb, ccm)
		}
	}
}

func Solve() {
	var n, m, a, b int
	g := map[int]map[int]bool{}
	gb := map[int]map[int]bool{}
	fmt.Scan(&n)
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		if _, ok := g[b]; !ok {
			g[b] = map[int]bool{}
		}
		g[b][a] = true

		if _, ok := g[a]; !ok {
			g[a] = map[int]bool{}
		}
		g[a][b] = true
	}
	visited := map[int]bool{}
	tin := make([]int, n+1)
	low := make([]int, n+1)
	timer := 0
	for i := 1; i <= n; i++ {
		if _, ok := visited[i]; !ok {
			timer = dfs(i, -1, timer, g, gb, visited, tin, low)
		}
	}
	bridges := 0
	for node := range gb {
		bridges += len(gb[node])
	}
	bridges /= 2
	// log.Println(bridges)
	vcc := map[int]int{}
	cc := 0
	for i := 1; i <= n; i++ {
		if _, ok := vcc[i]; !ok {
			cc++
			dfscc(i, cc, g, gb, vcc)
		}
	}
	gcb := map[int]map[int]bool{}
	for ba, nodes := range gb {
		for bb := range nodes {
			a := vcc[ba]
			b := vcc[bb]
			if _, ok := gcb[b]; !ok {
				gcb[b] = map[int]bool{}
			}
			gcb[b][a] = true

			if _, ok := gcb[a]; !ok {
				gcb[a] = map[int]bool{}
			}
			gcb[a][b] = true
		}
	}
	smaxdepth := -1
	v1 := map[int]bool{}
	v2 := map[int]bool{}
	for node := range gcb {
		maxnode := 0
		maxdepth := -1
		if _, ok := v1[node]; !ok {
			diadfs(node, -1, 0, &maxdepth, &maxnode, gcb, v1)
			// log.Println(maxnode)
			diadfs(maxnode, -1, 0, &maxdepth, &maxnode, gcb, v2)
		}
		if maxdepth > smaxdepth {
			smaxdepth = maxdepth
		}
	}
	// log.Println(smaxdepth)
	fmt.Println(bridges - smaxdepth)
}
