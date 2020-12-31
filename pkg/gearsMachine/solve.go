package gearsMachine

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
)

var bytes []byte
var l, max int

func fastScan(n *int) {
	b := bytes[l]
	*n = 1
	for b < 48 || b > 57 {
		if b == 0x2d {
			*n = -1
		}
		l++
		b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > max {
			*n *= result
			return
		}
		b = bytes[l]
	}
	*n *= result
}

func init() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
}

func bfs(node, cc int, g [][]int, side, dir, ccm []int) bool {
	fked := true
	q := list.New()
	dir[node] = 1
	for q.PushBack(node); q.Len() > 0; {
		next := q.Front()
		neigh := next.Value.(int)
		q.Remove(next)
		ccm[neigh] = cc
		neighdir := dir[neigh] * side[neigh] * -1
		for _, v := range g[neigh] {
			if dir[v] == 0 {
				dir[v] = neighdir * side[v]
				q.PushBack(v)
			} else {
				if dir[v] != neighdir*side[v] {
					fked = false
				}
			}
		}
	}
	return fked
}

func Solve() {
	var n, m, a, b, da, db, q int
	fastScan(&n)
	fastScan(&m)
	fastScan(&q)
	side := make([]int, n+1)
	dir := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fastScan(&side[i])
	}
	g := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		fastScan(&a)
		fastScan(&b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	cc := 0
	ccm := make([]int, n+1)
	ccs := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if dir[i] == 0 {
			cc++
			ccs[cc] = bfs(i, cc, g, side, dir, ccm)
		}
	}
	// log.Println(dir)
	for i := 1; i <= q; i++ {
		fastScan(&a)
		fastScan(&b)
		fastScan(&da)
		fastScan(&db)
		if !ccs[ccm[a]] || !ccs[ccm[b]] {
			fmt.Println("NO")
			continue
		}
		if ccm[a] != ccm[b] {
			fmt.Println("YES")
			continue
		}
		if da*db == dir[a]*dir[b] {
			fmt.Println("YES")
			continue
		}
		fmt.Println("NO")
	}
}
