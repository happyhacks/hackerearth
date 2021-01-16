package makingFriends

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var bytes []byte
var l, max int

func fastScan(n *int64) {
	b := bytes[l]
	*n = 1
	for b < 48 || b > 57 {
		if b == 0x2d {
			*n = -1
		}
		l++
		b = bytes[l]
	}

	result := int64(0)
	for 47 < b && b < 58 {
		result = (result << 3) + (result << 1) + int64(b-48)

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

func Solve() {
	var t, n, m int64
	fastScan(&t)
	for tst := int64(0); tst < t; tst++ {
		fastScan(&n)
		fastScan(&m)
		if m == 0 {
			fmt.Println("Yes")
			continue
		}
		if n%2 == 0 && m <= n/2 {
			fmt.Println("Yes")
			continue
		}
		fmt.Println("No")
	}
}
