package longestPath

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

func Solve() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
	var n, k, a, b int
	g := map[int]map[int]bool{}
	fastScan(&n)
	fastScan(&k)
	label := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fastScan(&label[i])
	}
	for i := 0; i < n-1; i++ {
		fastScan(&a)
		fastScan(&b)
		if label[a]%k != 0 || label[b]%k != 0 {
			continue
		}
		if _, ok := g[b]; !ok {
			g[b] = map[int]bool{}
		}
		g[b][a] = true

		if _, ok := g[a]; !ok {
			g[a] = map[int]bool{}
		}
		g[a][b] = true
	}
	smaxdepth := 0
	v1 := map[int]bool{}
	v2 := map[int]bool{}
	for node := range g {
		maxnode := 0
		maxdepth := 0
		if _, ok := v1[node]; !ok {
			diadfs(node, -1, 0, &maxdepth, &maxnode, g, v1)
			// log.Println(maxnode)
			diadfs(maxnode, -1, 0, &maxdepth, &maxnode, g, v2)
		}
		if maxdepth > smaxdepth {
			smaxdepth = maxdepth
		}
	}
	fmt.Println(smaxdepth)
}
