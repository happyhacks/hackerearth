package scc

import (
	"fmt"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func dfs1(src int, g map[int]map[int]bool, s *Stack, v map[int]bool) {
	if _, ok := v[src]; ok {
		return
	}
	v[src] = true
	for neigh := range g[src] {
		if _, ok := v[neigh]; ok {
			continue
		}
		dfs1(neigh, g, s, v)
	}
	s.Push(src)
}

func dfs2(src int, g map[int]map[int]bool, v map[int]bool) (x int) {
	if _, ok := v[src]; ok {
		return x
	}
	v[src] = true
	for neigh := range g[src] {
		if _, ok := v[neigh]; ok {
			continue
		}
		x += dfs2(neigh, g, v)
	}
	return x + 1
}

func Solve() {
	var n, m, a, b, score int
	g := map[int]map[int]bool{}
	fg := map[int]map[int]bool{}
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		if _, ok := g[a]; !ok {
			g[a] = map[int]bool{}
		}
		if _, ok := fg[b]; !ok {
			fg[b] = map[int]bool{}
		}
		g[a][b] = true
		fg[b][a] = true
	}
	visited := map[int]bool{}
	s := NewStack()
	for i := 1; i <= n; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		dfs1(i, g, s, visited)
	}
	visited = map[int]bool{}
	for s.Len() > 0 {
		src := s.Pop().(int)
		if _, ok := visited[src]; ok {
			continue
		}
		x := dfs2(src, fg, visited)
		if x%2 == 0 {
			score -= x
		} else {
			score += x
		}
	}
	fmt.Println(score)
}
