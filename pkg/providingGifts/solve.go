package providingGifts

import "fmt"

func Solve() {
	var n, t int64
	fmt.Scan(&t)
	for i := int64(0); i < t; i++ {
		fmt.Scan(&n)
		if n%2 == 1 {
			fmt.Println(1 + n/2)
		} else {
			fmt.Println(n / 2)
		}
	}
}
