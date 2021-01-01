package witchesHegwarts

import (
	"container/list"
	"fmt"
)

const Maxuint32 = ^uint32(0)

func min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func g(n uint32) uint32 {
	d := map[uint32]uint32{}
	q := list.New()
	d[n] = 0
	addr := func(c, dist uint32) bool {
		if _, ok := d[c]; !ok {
			d[c] = dist + 1
		} else {
			d[c] = min(d[c], dist+1)
		}
		q.PushBack(c)
		return c == 1
	}
	for q.PushBack(n); q.Len() > 0; {
		next := q.Front()
		neigh := next.Value.(uint32)
		if neigh == 1 {
			return d[neigh]
		}
		q.Remove(next)
		dist := d[neigh]
		if neigh%2 == 0 && addr(neigh/2, dist) {
			return d[neigh/2]
		}
		if neigh%3 == 0 && addr(neigh/3, dist) {
			return d[neigh/3]
		}
		if addr(neigh-1, dist) {
			return d[neigh-1]
		}

	}
	return 0
}

func Solve() {
	var n, a, i uint32
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&a)
		fmt.Println(g(a))
	}
}
