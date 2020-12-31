package hateEvenSubarrays

import (
	"fmt"
	"strconv"
	"strings"
)

type run struct {
	char   int
	length int
}

func rle(input string) []run {
	result := []run{}
	for len(input) > 0 {
		firstLetter := input[0]
		inputLength := len(input)
		input = strings.TrimLeft(input, string(firstLetter))
		if counter := inputLength - len(input); counter > 0 {
			c, _ := strconv.Atoi(string(firstLetter))
			r := run{char: c, length: counter}
			result = append(result, r)
		}
	}
	return result
}

func merge(runs []run) []run {
	result := []run{}
	if len(runs) == 0 {
		return result
	}
	last := runs[0]
	for i := 1; i < len(runs); i++ {
		if last.char == runs[i].char {
			last.length += runs[i].length
		} else {
			result = append(result, last)
			last = runs[i]
		}
	}
	result = append(result, last)
	return result
}

func minify(runs []run) []run {
	for i := 0; i < len(runs); i++ {
		runs[i].length %= 2
	}
	result := []run{}
	for i := 0; i < len(runs); i++ {
		if runs[i].length > 0 {
			result = append(result, runs[i])
		}
	}
	return result
}

func runLen(runs []run) (sum int) {
	for _, i := range runs {
		sum += i.length
	}
	return sum
}

func Solve() {
	var t int
	var s string
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&s)
		runs := rle(s)
		for runLen(runs) > len(runs) {
			runs = minify(runs)
			// log.Println(runs)

			runs = merge(runs)
			// log.Println(runs)

		}
		if len(runs) == 0 {
			fmt.Println("KHALI")
		} else {
			for _, i := range runs {
				for j := 0; j < i.length; j++ {
					fmt.Print(i.char)
				}
			}
			fmt.Println("")
		}
	}
}
