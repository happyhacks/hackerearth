package multipleSubtrees

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var bytes []byte
var l, max int

func fastScan(n *int) {
	b := bytes[l]

	for b < 48 || b > 57 {
		l++
		b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > max {
			*n = result
			return
		}
		b = bytes[l]
	}
	*n = result
}

func dfs(node, parent, cSum int, g map[int]map[int]bool, v []bool, p, s []int) {
	if v[node] {
		return
	}
	v[node] = true
	if _, ok := g[node][parent]; !ok {
		s[node] = len(g[node])
	} else {
		s[node] = len(g[node]) - 1
	}
	p[node] = cSum
	for neigh := range g[node] {
		if neigh == parent {
			continue
		}
		dfs(neigh, node, p[node]+s[node]-1, g, v, p, s)
	}
}

func Solve() {
	var n, m, a, b int
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
	// ws := bufio.NewWriter(os.Stdout)
	g := map[int]map[int]bool{}
	fastScan(&n)
	for i := 0; i < n-1; i++ {
		fastScan(&a)
		fastScan(&b)
		if _, ok := g[b]; !ok {
			g[b] = map[int]bool{}
		}
		g[b][a] = true

		if _, ok := g[a]; !ok {
			g[a] = map[int]bool{}
		}
		g[a][b] = true
	}
	visited := make([]bool, n+1)
	parentSum := make([]int, n+1)
	selfChild := make([]int, n+1)
	dfs(1, -1, 0, g, visited, parentSum, selfChild)
	// log.Println(parentSum, selfChild)
	fastScan(&m)
	for i := 0; i < m; i++ {
		fastScan(&a)
		fmt.Println(parentSum[a] + selfChild[a])
	}
}
