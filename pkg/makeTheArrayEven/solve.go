package makeTheArrayEven

import "fmt"

func solve() {
	var n, a int
	fmt.Scan(&n)
	sum := 0
	bits := make([]int8, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		bits[i] = int8(a % 2)
		if i != 0 && bits[i] == 1 {
			if bits[i-1] == 1 {
				sum++
				bits[i], bits[i-1] = 0, 0
			}
		}
	}
	for i := 0; i < n-1; i++ {
		a, b := bits[i], bits[i+1]
		if a == 0 && b == 0 {
			i++
		} else if a == 1 && b == 1 {
			sum++
			bits[i], bits[i+1] = 0, 0
			i++
		} else {
			sum += 2
			bits[i], bits[i+1] = 0, 0
			i++
		}
	}
	if bits[n-1] == 1 {
		sum += 2
	}
	fmt.Println(sum)
}

func Solve() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
