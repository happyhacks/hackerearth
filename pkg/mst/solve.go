package mst

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	a, b, weight int64
}

type Edges []Edge

func (h Edges) Len() int {
	return len(h)
}

func (h Edges) Less(i, j int) bool {
	if h[i].weight == h[j].weight {
		if h[i].a == h[j].a {
			return h[i].b < h[j].b
		}
		return h[i].a < h[j].a
	}
	return h[i].weight < h[j].weight
}

func (h Edges) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Edges) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *Edges) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Element struct {
	parent *Element
	rank   int64
	Data   interface{}
}

func NewElement() *Element {
	s := &Element{}
	s.parent = s
	return s
}

func (e *Element) Find() *Element {
	for e.parent != e {
		e.parent = e.parent.parent
		e = e.parent
	}
	return e
}

func Union(e1, e2 *Element) {
	e1Root := e1.Find()
	e2Root := e2.Find()
	if e1Root == e2Root {
		return
	}

	switch {
	case e1Root.rank < e2Root.rank:
		e1Root.parent = e2Root
	case e1Root.rank > e2Root.rank:
		e2Root.parent = e1Root
	default:
		e2Root.parent = e1Root
		e1Root.rank++
	}
}

func Solve() {
	var m, n, a, b, w, i int64
	fmt.Scanf("%d %d", &n, &m)
	nodes := make([]*Element, n+1)
	g := &Edges{}
	heap.Init(g)
	for i = 0; i <= n; i++ {
		nodes[i] = NewElement()
	}

	for i = 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&w)
		heap.Push(g, Edge{a, b, w})
	}
	w = 0
	for i = 0; i < m; i++ {
		cheapest := heap.Pop(g).(Edge)
		if nodes[cheapest.a].Find() == nodes[cheapest.b].Find() {
			continue
		}
		Union(nodes[cheapest.a], nodes[cheapest.b])
		w += cheapest.weight
	}
	fmt.Println(w)
}
