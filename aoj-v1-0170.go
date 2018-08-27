package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var f = make([]string, 10)
var s, w = make([]int, 10), make([]int, 10)
var n int

func main() {
	for scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			tmpf := buf[0]
			tmpw, _ := strconv.Atoi(buf[1])
			tmps, _ := strconv.Atoi(buf[2])
			f[i] = tmpf
			s[i] = tmps
			w[i] = tmpw
		}
		sort()
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if w[j] < w[j+1] {
					swap(j)
					a := get_weight(j + 1)
					if a > s[j] {
						swap(j)
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			fmt.Println(f[i])
		}
	}
}

func sort() {
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] < s[j+1] {
				swap(j)
			} else if s[j] == s[j+1] && w[j] < w[j+1] {
				swap(j)
			}
		}
	}
}

func swap(x int) {
	w[x], w[x+1] = w[x+1], w[x]
	s[x], s[x+1] = s[x+1], s[x]
	f[x], f[x+1] = f[x+1], f[x]
}

func get_weight(x int) int {
	ret := 0
	for i := x; i < n; i++ {
		ret += w[i]
	}
	return ret
}
