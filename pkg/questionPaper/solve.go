package questionPaper

import (
	"fmt"
)

// type

func bfs(n, a, b int, d map[int]int) {
	q := []int{}
	d[0] = 0
	for q = append(q, 0); len(q) > 0; {
		node := q[0]
		q = q[1:]
		if d[node] >= n {
			return
		}
		if _, ok := d[node+a]; !ok {
			d[node+a] = d[node] + 1
			q = append(q, node+a)
		}
		if _, ok := d[node-b]; !ok {
			d[node-b] = d[node] + 1
			q = append(q, node-b)
		}
	}
}

func Solve() {
	var t, n, a, b int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		d := make(map[int]int)
		fmt.Scan(&n)
		fmt.Scan(&a)
		fmt.Scan(&b)
		bfs(n, a, b, d)
		fmt.Println(len(d))
	}
}
