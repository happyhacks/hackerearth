package arrReverse

import "fmt"

func postOrderPrint(depth int) {
	if depth <= 0 {
		return
	}
	var curr int
	fmt.Scan(&curr)
	postOrderPrint(depth - 1)
	fmt.Println(curr)
}

func Solve() {
	var depth int
	fmt.Scan(&depth)
	postOrderPrint(depth)
}
