package visitingFriends

import (
	"bufio"
	"fmt"

	"io/ioutil"
	"os"
	"sort"
)

var bytes []byte
var l, max int

func fastScan(n *int) {
	b := bytes[l]

	for b < 48 || b > 57 {
		l++
		b = bytes[l]
	}

	result := 0
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int(b-48)

		l++
		if l > max {
			*n = result
			return
		}
		b = bytes[l]
	}
	*n = result
}

func init() {
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
}

var fact = []uint64{1, 0, 1, 2, 9, 44, 265, 1854, 14833, 133496, 1334961, 14684570, 176214841, 2290792932, 32071101049, 481066515734, 7697064251745, 130850092279664, 2355301661033953, 44750731559645106, 895014631192902121}

func dfscc(node, cc int, g [][]int, ccm, ccc map[int]int) {
	if _, ok := ccm[node]; ok {
		return
	}
	ccm[node] = cc
	ccc[cc]++
	for _, neigh := range g[node] {
		if _, ok := ccm[neigh]; !ok {
			dfscc(neigh, cc, g, ccm, ccc)
		}
	}
}

func solve() {
	var n, m, a, b, i int
	var cc int
	fastScan(&n)
	fastScan(&m)
	g := make([][]int, n+1)
	for i = 0; i < m; i++ {
		fastScan(&a)
		fastScan(&b)
		g[b] = append(g[b], a)
		g[a] = append(g[a], b)
	}
	// visited := map[int]bool{}
	cc = 0
	vcc := map[int]int{}
	ccc := map[int]int{}
	for i = 1; i <= n; i++ {
		if _, ok := vcc[i]; !ok {
			cc++
			dfscc(i, cc, g, vcc, ccc)
		}
	}
	// log.Pruintln(vcc, ccc)
	ws := bufio.NewWriter(os.Stdout)

	// fmt.Println(len(ccc))
	ws.WriteString(fmt.Sprint(len(ccc)) + "\n")
	counts := make([]int, 0, len(ccc))
	for _, count := range ccc {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	for i := len(ccc) - 1; i >= 0; i-- {
		// fmt.Printf("%d ", count)
		ws.WriteString(fmt.Sprint(fact[counts[i]]) + " ")
	}
	ws.WriteString("\n")
	ws.Flush()
}

func Solve() {
	var t int
	fastScan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
