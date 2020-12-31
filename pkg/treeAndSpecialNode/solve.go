package treeAndSpecialNode

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

func init() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
}

func dfs(node, parent int, g [][]int, c []int, p, s map[int]bool, count *int) {
	// log.Println(node, parent)
	if seen, ok := p[c[node]]; ok && seen {
		return
	}
	s[node] = true
	*count++
	p[c[node]] = true
	for _, neigh := range g[node] {
		if neigh == parent {
			continue
		}
		dfs(neigh, node, g, c, p, s, count)
	}
	p[c[node]] = false
}

func Solve() {
	var n, a, b, i int
	fastScan(&n)
	g := make([][]int, n+1)
	c := make([]int, n+1)
	for i = 1; i <= n; i++ {
		fastScan(&c[i])
	}
	for i = 0; i < n-1; i++ {
		fastScan(&a)
		fastScan(&b)
		g[b] = append(g[b], a)
		g[a] = append(g[a], b)
	}
	rootPath := map[int]bool{}
	special := map[int]bool{}
	var count int
	dfs(1, -1, g, c, rootPath, special, &count)
	fmt.Println(count)
}
