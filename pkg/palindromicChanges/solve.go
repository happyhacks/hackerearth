package palindromicChanges

import (
	"fmt"
	"math"
)

type edge struct {
	to int
	wt float64
}

func floydWarshall(g [][]edge) [][]float64 {
	dist := make([][]float64, len(g))
	for i := range dist {
		di := make([]float64, len(g))
		for j := range di {
			di[j] = math.Inf(1)
		}
		di[i] = 0
		dist[i] = di
	}
	for u, graphs := range g {
		for _, v := range graphs {
			dist[u][v.to] = v.wt
		}
	}
	for k, dk := range dist {
		for _, di := range dist {
			for j, dij := range di {
				if d := di[k] + dk[j]; dij > d {
					di[j] = d
				}
			}
		}
	}
	return dist
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Solve() {
	var s string
	fmt.Scan(&s)
	var m int
	fmt.Scan(&m)
	var a, b string
	var w float64
	g := make([][]edge, 26)
	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&w)
		// log.Println(a, b, w)
		g[a[0]-'a'] = append(g[a[0]-'a'], edge{int(b[0] - 'a'), w})
	}
	d := floydWarshall(g)
	// log.Println(d)
	ls := len(s)
	startidx := len(s) / 2
	if ls%2 == 1 {
		startidx++
	}
	sm := int64(0)
	for i := startidx; i < ls; i++ {
		a := int(s[i] - 'a')
		b := int(s[ls-i-1] - 'a')
		// log.Println(string(byte(a+'a')), string(byte(b+'a')), d[a][b], d[b][a])
		if a == b {
			continue
		}
		// if d[a][b] > d[b][a] {
		// 	sm += int64(d[b][a])
		// } else {
		// 	sm += int64(d[a][b])
		// }
		pairmin := min(d[a][b], d[b][a])
		for intermediate := 0; intermediate < 26; intermediate++ {
			pairmin = min(pairmin, d[a][intermediate]+d[b][intermediate])
		}
		sm += int64(pairmin)
	}
	fmt.Println(int64(sm))
}
