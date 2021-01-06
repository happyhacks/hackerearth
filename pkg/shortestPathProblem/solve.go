package shortestPathProblem

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/ioutil"
	"os"
)

const MaxUint = ^uint64(0)

var bytes []byte
var l, max int

func fastScan(n *int) {
	if l >= len(bytes) {
		*n = 0
		return
	}
	b := bytes[l]
	*n = 1
	for b < 48 || b > 57 {
		if b == 0x2d {
			*n = -1
		}
		l++
		if l >= len(bytes) {
			*n = 0
			return
		}
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

type Item struct {
	node int
	dist uint64
	idx  int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.idx = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.idx = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}

func dijkstra(src int, g [][]int, wt map[int]map[int]int, d []uint64, v []bool) {
	d[src] = 0
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	q.Push(&Item{node: src, dist: 0})
	for q.Len() > 0 {
		next := heap.Pop(&q).(*Item)
		nextNode := next.node
		// log.Println(nextNode)
		if v[nextNode] {
			continue
		}
		v[nextNode] = true
		for _, neigh := range g[nextNode] {
			weight := wt[nextNode][neigh]
			if d[neigh] > d[nextNode]+uint64(weight) {
				d[neigh] = d[nextNode] + uint64(weight)
				heap.Push(&q, &Item{node: neigh, dist: d[neigh]})
			}
		}
	}
}

func Solve() {
	var n, m, a, b, w int
	fastScan(&n)
	fastScan(&m)
	g := make([][]int, n+1)
	v := make([]bool, n+1)
	wt := make(map[int]map[int]int)
	for i := 0; i < m; i++ {
		fastScan(&a)
		fastScan(&b)
		fastScan(&w)
		g[a] = append(g[a], b)
		if _, ok := wt[a]; !ok {
			wt[a] = make(map[int]int)
		}
		if oldwt, ok := wt[a][b]; !ok || oldwt > w {
			wt[a][b] = w
		}
	}
	distf := make([]uint64, n+1)
	for i := 0; i <= n; i++ {
		distf[i] = MaxUint
		v[i] = false
	}
	dijkstra(1, g, wt, distf, v)
	for i := 2; i <= n; i++ {
		fmt.Printf("%d ", distf[i])
	}
}
