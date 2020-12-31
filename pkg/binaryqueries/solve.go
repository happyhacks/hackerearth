package binaryqueries

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
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

func Solve() {
	var n, q, L, R, t int
	bytes, _ = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	max = len(bytes) - 1
	ws := bufio.NewWriter(os.Stdout)
	fastScan(&n)
	fastScan(&q)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fastScan(&arr[i])
	}
	for i := 1; i <= q; i++ {
		fastScan(&t)
		if t == 0 {
			fastScan(&L)
			fastScan(&R)
			if arr[R] == 0 {
				ws.WriteString("EVEN\n")
			} else if arr[R] == 1 {
				ws.WriteString("ODD\n")
			} else {
				log.Println(arr[R])
				panic("err")
			}
		} else {
			fastScan(&t)
			if arr[t] == 0 {
				arr[t] = 1
			} else {
				arr[t] = 0
			}
		}
	}
	ws.Flush()
}
