package crossTheRiver

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const MaxUint = ^uint64(0)
const MaxInt = int64(MaxUint >> 1)

var bytes []byte
var l, max int

func fastScan(n *int64) {
	b := bytes[l]
	*n = 1
	for b < 48 || b > 57 {
		if b == 0x2d {
			*n = -1
		}
		l++
		b = bytes[l]
	}

	result := int64(0)
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int64(b-48)

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

type point struct {
	x, y int64
}

type rock struct {
	p point
	r int64
}

func dist(a, b point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func bfs(src int64, g [][]int64, d []int64) {
	q := list.New()
	d[src] = 0
	for q.PushBack(src); q.Len() > 0; {
		next := q.Front()
		node := next.Value.(int64)
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
	var a, b, n int64
	fastScan(&n)
	rocks := make([]rock, n+2)
	g := make([][]int64, n+2)
	for i := int64(0); i < n; i++ {
		fastScan(&rocks[i+1].p.x)
		fastScan(&rocks[i+1].p.y)
		fastScan(&rocks[i+1].r)
	}
	fastScan(&a)
	fastScan(&b)
	for i := int64(1); i <= n; i++ {
		if dist(rocks[i].p, point{rocks[i].p.x, b}) <= float64(rocks[i].r) {
			g[0] = append(g[0], i)
			g[i] = append(g[i], 0)
		}
		if dist(rocks[i].p, point{rocks[i].p.x, a}) <= float64(rocks[i].r) {
			g[n+1] = append(g[n+1], i)
			g[i] = append(g[i], n+1)
		}
		for j := int64(1); j <= n; j++ {
			if dist(rocks[i].p, rocks[j].p) <= float64(rocks[i].r+rocks[j].r) {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}
	dist := make([]int64, n+2)
	for i := int64(0); i < n+2; i++ {
		dist[i] = MaxInt
	}
	bfs(0, g, dist)
	if dist[n+1] == MaxInt {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[n+1] - 1)
	}
}

func Solve() {
	var t int64
	fastScan(&t)
	for i := int64(0); i < t; i++ {
		solve()
	}
}
