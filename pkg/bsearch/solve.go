package bsearch

import (
	"fmt"
	"sort"
)

func Solve() {
	var n, a, q int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&a)
		fmt.Println(sort.SearchInts(arr, a) + 1)
	}
}
