package connectedHorses

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

const mod = uint64(10e8 + 7)

type point struct {
	x, y int
}

func (p *point) horseNeigh(pointsAll map[point]bool) []*point {
	dx := []int{-2, -2, 1, -1, 1, -1, 2, 2}
	dy := []int{-1, 1, 2, 2, -2, -2, -1, 1}
	points := make([]*point, 0)
	for i := 0; i < 8; i++ {
		pt := &point{p.x + dx[i], p.y + dy[i]}
		if _, ok := pointsAll[*pt]; ok {
			points = append(points, pt)
		}
	}
	return points
}

type grid struct {
	X, Y int
	V    map[point]int
	P    map[point]bool
	C    map[int]int
}

func newGrid(n, m int, p map[point]bool) *grid {
	g := grid{X: m, Y: n, P: p}
	g.V = make(map[point]int)
	g.C = make(map[int]int)
	return &g
}

func (g *grid) inBounds(p *point) bool {
	return p.x >= 0 && p.x < g.X && p.y >= 0 && p.y < g.Y
}

func (g *grid) dfs(p *point, c int) {
	if _, ok := g.V[*p]; ok {
		return
	}
	if !g.inBounds(p) {
		return
	}
	if _, ok := g.P[*p]; !ok {
		return
	}
	g.V[*p] = c
	g.C[c]++
	for _, i := range p.horseNeigh(g.P) {
		g.dfs(i, c)
	}
}

func modFactorial(n int, modulus uint64) uint64 {
	result := uint64(1)
	for i := uint64(2); i <= uint64(n); i++ {
		result = (result * i) % modulus
	}
	return result % modulus
}

func solve() {
	var n, m, q, a, b int
	fastScan(&n)
	fastScan(&m)
	fastScan(&q)
	p := map[point]bool{}
	for i := 0; i < q; i++ {
		fastScan(&a)
		fastScan(&b)
		p[point{a - 1, b - 1}] = true
	}
	g := newGrid(m, n, p)
	c := 0
	for i := range p {
		if g.V[i] == 0 {
			c++
			g.dfs(&i, c)
		}
	}
	// log.Println(g.C)
	s := uint64(1)
	for _, count := range g.C {
		if count > 1 {
			s *= modFactorial(count, mod)
			s %= mod
		}
	}
	fmt.Println(s)
}

func Solve() {
	var t int
	fastScan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
