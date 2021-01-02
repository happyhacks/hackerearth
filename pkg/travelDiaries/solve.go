package travelDiaries

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
)

type point struct {
	x, y int
}

func (p *point) right() *point {
	return &point{p.x + 1, p.y}
}

func (p *point) left() *point {
	return &point{p.x - 1, p.y}
}

func (p *point) up() *point {
	return &point{p.x, p.y - 1}
}

func (p *point) down() *point {
	return &point{p.x, p.y + 1}
}

type grid struct {
	X, Y int
	G, D [][]int
}

func newGrid(n, m int) *grid {
	g := grid{X: m, Y: n, G: make([][]int, n), D: make([][]int, n)}
	for i := 0; i < n; i++ {
		g.G[i] = make([]int, m)
		g.D[i] = make([]int, m)
	}
	return &g
}

func (g *grid) inBounds(p *point) bool {
	return p.x >= 0 && p.x < g.X && p.y >= 0 && p.y < g.Y
}

func (g *grid) at(p *point) int {
	return g.G[p.y][p.x]
}

func (g *grid) setDist(p *point, d int) {
	g.D[p.y][p.x] = d
}

func (g *grid) getDist(p *point) int {
	return g.D[p.y][p.x]
}

func (g *grid) bfs() {
	q := list.New()
	for i := 0; i < g.Y; i++ {
		for j := 0; j < g.X; j++ {
			if g.G[i][j] == 2 {
				q.PushBack(&point{j, i})
			}
		}
	}
	checkAppend := func(pt *point, d int) {
		if !g.inBounds(pt) {
			return
		}
		if g.at(pt) != 1 {
			return
		}
		ptd := g.getDist(pt)
		if ptd != 0 {
			if ptd > d+1 {
				g.setDist(pt, d+1)
				q.PushBack(pt)
			}
		} else {
			g.setDist(pt, d+1)
			q.PushBack(pt)
		}
	}
	for q.Len() > 0 {
		next := q.Front()
		neigh := next.Value.(*point)
		q.Remove(next)
		d := g.getDist(neigh)
		// log.Println(neigh, d)
		checkAppend(neigh.down(), d)
		checkAppend(neigh.up(), d)
		checkAppend(neigh.left(), d)
		checkAppend(neigh.right(), d)
	}
}

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

func Solve() {
	var n, m int
	fastScan(&n)
	fastScan(&m)
	g := newGrid(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fastScan(&g.G[i][j])
		}
	}
	g.bfs()
	mx := 0
	var single bool
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g.D[i][j] > mx {
				mx = g.D[i][j]
			}
			if g.G[i][j] == 1 && g.D[i][j] == 0 {
				single = true
			}
		}
	}
	if !single {
		fmt.Println(mx)
	} else {
		fmt.Println(-1)
	}
}
