package bennyAndTheUniverse

import (
	"container/list"
	"fmt"
)

const MaxUint = ^uint64(0)

func bfs(d []uint64, m []uint64, md uint64) {
	v := map[uint64]uint64{}
	q := list.New()
	q.PushBack(uint64(0))
	v[0] = 0
	for q.Len() > 0 {
		nxt := q.Front()
		node := nxt.Value.(uint64)
		q.Remove(nxt)
		// log.Println(node)
		if val, ok := v[node]; ok {
			if val > 10e8+10 {
				continue
			}
		}
		for _, n := range d {
			next := n + v[node]
			nextMod := next % md
			// log.Println(next, nextMod)
			if m[nextMod] > next {
				m[nextMod] = next
				q.PushBack(nextMod)
				v[nextMod] = next
			}
		}
	}
}

func Solve() {
	var n, q, a uint64
	fmt.Scan(&n)
	fmt.Scan(&q)
	d := make([]uint64, n+1)
	mind := MaxUint
	// minidx := uint64(0)
	for i := uint64(0); i < n; i++ {
		fmt.Scan(&d[i])
		if d[i] < mind {
			mind = d[i]
			// minidx = i
		}
	}
	mods := make([]uint64, mind)
	for i := uint64(1); i < mind; i++ {
		mods[i] = MaxUint
	}
	mods[0] = 0
	bfs(d, mods, mind)
	// log.Println(mods)
	for i := uint64(0); i < q; i++ {
		fmt.Scan(&a)
		md := a % mind
		if mods[md] <= a {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
