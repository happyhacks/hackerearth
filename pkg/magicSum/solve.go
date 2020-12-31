package magicSum

import (
	"fmt"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

func lca(a, b int) int {
	for {
		if a == b {
			return a
		}
		a /= 2
		b /= 2
	}
}

func solve() {
	var n int
	fmt.Scan(&n)
	val := make([]int, n+1)
	acval := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&val[i])
		acval[i] = val[i]
	}
	// dist := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		// dist[i] = make([]int, n+1)
		acval[i] += acval[i/2]
	}
	maxd := MinInt
	for i := (n + 1) / 2; i <= n; i++ {
		if val[i] > maxd {
			maxd = val[i]
		}
		for j := i + 1; j <= n; j++ {
			l := lca(i, j)
			d := acval[i] + acval[j] - 2*acval[l] + val[l]
			if d > maxd {
				maxd = d
			}
			// dist[i][j] = d
			// log.Println(i, j, d)
		}
	}
	fmt.Println(maxd)
}

func Solve() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
