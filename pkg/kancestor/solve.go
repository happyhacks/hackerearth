package kancestor

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
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

const maxn = 1000001

func dfs(g map[int]map[int]bool, color []int, node, par, k int, colors [][]int, solutions map[int]int) {
	c := color[node]
	tree := colors[c]
	l := len(tree)
	if l >= k {
		solutions[node] = tree[l-k]
	}
	for neigh := range g[node] {
		if neigh != par {
			colors[c] = append(colors[c], node)
			dfs(g, color, neigh, node, k, colors, solutions)
			colors[c] = colors[c][:len(colors[c])-1]
		}
	}
}

func Solve() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
	var n, k, a, b int
	g := map[int]map[int]bool{}
	fastScan(&n)
	fastScan(&k)
	color := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fastScan(&a)
		color[i] = a
	}
	for i := 1; i < n; i++ {
		fastScan(&a)
		fastScan(&b)
		if _, ok := g[b]; !ok {
			g[b] = map[int]bool{}
		}
		if _, ok := g[a]; !ok {
			g[a] = map[int]bool{}
		}
		g[a][b] = true
		g[b][a] = true
	}
	solutions := map[int]int{}
	colors := make([][]int, maxn)

	dfs(g, color, 1, -1, k, colors, solutions)
	ws := bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		if s, ok := solutions[i]; ok {
			ws.WriteString(strconv.Itoa(s) + " ")
		} else {
			ws.WriteString("-1 ")
		}
	}
	ws.Flush()
}
