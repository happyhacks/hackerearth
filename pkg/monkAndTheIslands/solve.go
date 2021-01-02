package monkAndTheIslands

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)

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

func bfs(src int, g [][]int, d []int) {
	q := list.New()
	d[src] = 0
	for q.PushBack(src); q.Len() > 0; {
		next := q.Front()
		node := next.Value.(int)
		// log.Println(node)
		q.Remove(next)
		for _, neigh := range g[node] {
			if d[neigh] > d[node]+1 {
				d[neigh] = d[node] + 1
				q.PushBack(neigh)
			}
		}
	}
}

func solve() {
	var n, m, a, b int
	fastScan(&n)
	fastScan(&m)
	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		fastScan(&a)
		fastScan(&b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		d[i] = MaxInt
	}
	bfs(1, g, d)
	fmt.Println(d[n])
}

func Solve() {
	var t int
	fastScan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
