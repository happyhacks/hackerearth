package escapeFromGrid

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

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
			if g.G[i][j] == -1 {
				q.PushBack(&point{j, i})
			}
		}
	}
	checkAppend := func(pt *point, d int) {
		if !g.inBounds(pt) {
			return
		}
		if g.at(pt) != 0 {
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
	defer writer.Flush()
	var n, m int
	var pt *point
	fmt.Scan(&n)
	fmt.Scan(&m)
	g := newGrid(n+2, m+2)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&g.G[i+1][j+1])
			if g.G[i+1][j+1] == 2 {
				pt = &point{j + 1, i + 1}
			}
			if g.G[i+1][j+1] == 0 {
				g.D[i+1][j+1] = int(1e8)
			}
		}
	}
	for _, i := range []int{0, n + 1} {
		for j := 0; j < m+2; j++ {
			g.G[i][j] = -1
		}
	}
	for _, j := range []int{0, m + 1} {
		for i := 0; i < n+2; i++ {
			g.G[i][j] = -1
		}
	}
	// log.Println("grid")
	// for i := 0; i < n+2; i++ {
	// 	for j := 0; j < m+2; j++ {
	// 		fmt.Print(g.G[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }
	g.bfs()
	// log.Println("dist")
	md := int(1e8)
	// for i := 0; i < n+2; i++ {
	// 	for j := 0; j < m+2; j++ {
	// 		// if g.G[i][j] == 2 {
	// 		// 	p := point{j, i}
	// 		// 	log.Println(g.getDist(p.down()))
	// 		// 	log.Println(g.getDist(p.up()))
	// 		// 	log.Println(g.getDist(p.right()))
	// 		// 	log.Println(g.getDist(p.left()))
	// 		// }
	// 		fmt.Print(g.D[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }
	for _, p := range []*point{pt.down(), pt.up(), pt.left(), pt.right()} {
		if g.at(p) == 0 {
			d := g.getDist(p)
			if md > d {
				md = d
			}
		} else if g.at(p) == -1 {
			md = 0
		}

	}
	if md == int(1e8) {
		md = -1
	}
	fmt.Println(md)
}
