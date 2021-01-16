package minimumCostPathInAGrid

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const maxG = 100001

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func logarr(a []int) {
	for idx, val := range a {
		fmt.Printf("%d:%d ", idx, val)
	}
	fmt.Println()
}

func Solve() {
	cost := make([]int, maxG)
	lg := make([]float64, maxG)
	for i := 1; i < maxG; i++ {
		lg[i] = math.Log(float64(i))
	}
	for j := 2; j < maxG; j++ {
		if cost[j] != 0 {
			continue
		}
		for i := j; i < maxG; i += j {
			pow := 0
			div := j
			for i%div == 0 {
				// log.Println(i, div)
				pow++
				div *= j
			}
			cost[i] += int(pow)
		}
	}
	// logarr(cost)
	var t, n, m, c int
	fastScan(&t)
	for tst := 0; tst < t; tst++ {
		fastScan(&n)
		fastScan(&m)
		g := make([][]int, n)
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			g[i] = make([]int, m)
			dp[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fastScan(&c)
				// log.Println(c, cost[c])
				g[i][j] = cost[c]
				if i == 0 && j == 0 {
					dp[i][j] = g[i][j]
				}
				if i == 0 && j > 0 {
					dp[i][j] = g[i][j] + dp[i][j-1]
				}
				if j == 0 && i > 0 {
					dp[i][j] = g[i][j] + dp[i-1][j]
				}
			}
		}
		// log.Println(g)
		// log.Println(dp)
		for i := 1; i < n; i++ {
			for j := 1; j < m; j++ {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + g[i][j]
			}
		}
		fmt.Println(dp[n-1][m-1])
	}
}
