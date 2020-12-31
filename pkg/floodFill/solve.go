package floodFill

import "fmt"

func dfs(x, y, n, m int, g [][]int) {
	if x < 0 || x >= n || y < 0 || y >= m {
		return
	}
	if g[x][y] != 1 {
		return
	}
	g[x][y] = -1
	dfs(x-1, y, n, m, g)
	dfs(x+1, y, n, m, g)
	dfs(x, y-1, n, m, g)
	dfs(x, y+1, n, m, g)
}

func Solve() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&g[i][j])
		}
	}
	dfs(0, 0, n, m, g)
	if g[n-1][m-1] == -1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
