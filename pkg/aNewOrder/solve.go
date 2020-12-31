package aNewOrder

import (
	"fmt"
	"log"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func firstDiff(a, b string) (int, int, error) {
	ml := min(len(a), len(b))
	for i := 0; i < ml; i++ {
		if a[i] != b[i] {
			return int(a[i] - 'a'), int(b[i] - 'a'), nil
		}
	}
	return 0, 0, fmt.Errorf("same strings")
}

func Solve() {
	var n int
	fmt.Scan(&n)
	chars := map[int]bool{}
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		for _, c := range arr[i] {
			chars[int(c)-'a'] = true
		}
	}
	g := make([][]bool, 26)
	for i := 0; i < 26; i++ {
		g[i] = make([]bool, 26)
	}
	indeg := make([]int, 26)
	for i := 0; i < n-1; i++ {
		a, b, err := firstDiff(arr[i], arr[i+1])
		log.Printf("%c->%c\n", 'a'+a, 'a'+b)
		if err != nil {
			continue
		}
		if !g[a][b] {
			indeg[b]++
		}
		g[a][b] = true
	}
	log.Println(indeg)
	visited := make([]bool, 26)
	ranks := make([][]int, 26)
	for ctr := 0; ; ctr++ {
		found := []int{}
		for i := 0; i < 26; i++ {
			if indeg[i] == 0 && chars[i] && !visited[i] {
				log.Println(i)
				visited[i] = true
				found = append(found, i)
				ranks[ctr] = append(ranks[ctr], i)
			}
		}
		if len(found) == 0 {
			break
		}
		for _, i := range found {
			for idx, freed := range g[i] {
				if freed {
					indeg[idx]--
					g[i][idx] = false
				}
			}
		}
	}
	// log.Println(ranks)
	for _, level := range ranks {
		for _, b := range level {
			fmt.Print(string('a' + b))
		}
		if len(level) > 0 {
			fmt.Println()
		}
	}
}
