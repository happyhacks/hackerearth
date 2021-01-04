package bracketSequence

import (
	"fmt"
	"log"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Solve() {
	var s string
	fmt.Scan(&s)
	mp := map[int]int{}
	var o, c int
	for _, i := range s {
		mp[o-c]++
		if i == '(' {
			o++
		} else if i == ')' {
			c++
		}
	}
	log.Println(mp)
	if o == c {
		min := MaxInt
		for key := range mp {
			if key < min {
				min = key
			}
		}
		fmt.Println(mp[min])
	} else {
		fmt.Println(0)
	}
}
