package robotInGrid

import (
	"bufio"
	"container/list"
	"fmt"
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

func Solve() {
	var n, m, q, sx, sy int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q)
	g := newGrid(n, m)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		for j, c := range s {
			if c == '*' {
				g.G[i][j] = -1
			} else if c == 'O' {
				g.G[i][j] = 1
			}
		}
	}
	fmt.Scan(&sy)
	fmt.Scan(&sx)
	g.G[sy-1][sx-1] = 2
	g.bfs()
	// log.Println(g.D)
	for i := 0; i < q; i++ {
		fmt.Scan(&sy)
		fmt.Scan(&sx)
		sy--
		sx--
		if g.G[sy][sx] == 1 && g.D[sy][sx] != 0 {
			fmt.Println(g.D[sy][sx])
		} else {
			fmt.Println(-1)
		}
	}
}
