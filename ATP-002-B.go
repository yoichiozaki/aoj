package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	n, _ := strconv.ParseInt(buf[0], 10, 64)
	m, _ := strconv.ParseInt(buf[1], 10, 64)
	p, _ := strconv.ParseInt(buf[2], 10, 64)
	fmt.Println(pow_mod(n, p, m))
}

func pow_mod(n, p, m int64) int64 {
	var r int64
	r = 1
	for ; p > 0; p >>= 1 {
		if p&1 != 0 {
			r = (r * n) % m
		}
		n = (n * n) % m
	}
	return r
}
