package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	MAX  = 1000000
	SQRT = 1000
)

func main() {
	comp := make([]int, MAX+2)
	sieve := func() {
		for i := 3; i < SQRT; i += 2 {
			if comp[i] == 0 {
				for j := i * i; j < MAX; j += i {
					comp[j] = 1
				}
			}
		}
	}
	sieve()
	tbl := make([]int, MAX+2)
	tbl[2] = 1
	k := 1
	for i := 3; i < MAX; i += 2 {
		if comp[i] == 0 {
			k++
		}
		tbl[i+1], tbl[i] = k, k
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		ans := 0
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			p, _ := strconv.Atoi(buf[0])
			m, _ := strconv.Atoi(buf[1])
			a, b := p-m, p+m
			if a < 2 {
				a = 2
			}
			if b > MAX {
				b = MAX
			}
			ans += tbl[b] - tbl[a-1] - 1
		}
		if ans < 0 {
			ans = 0
		}
		fmt.Println(ans)
	}
}
