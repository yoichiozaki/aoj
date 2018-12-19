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
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		s, p, q, r := make([]int, n), make([]int, n), make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			s[i], _ = strconv.Atoi(buf[0])
			p[i], _ = strconv.Atoi(buf[1])
			q[i], _ = strconv.Atoi(buf[2])
			r[i], _ = strconv.Atoi(buf[3])
		}
		scanner.Scan()
		buf_ := strings.Split(scanner.Text(), " ")
		P, _ := strconv.Atoi(buf_[0])
		Q, _ := strconv.Atoi(buf_[1])
		R, _ := strconv.Atoi(buf_[2])
		C, _ := strconv.Atoi(buf_[3])
		flag := false
		for i := 0; i < n; i++ {
			c := p[i]<<2 + r[i]<<2 + q[i]<<3 + q[i]
			if p[i] > P || q[i] > Q || r[i] > R || c > C {
				continue
			} else {
				fmt.Println(s[i])
				flag = true
			}
		}
		if !flag {
			fmt.Println("NA")
		}
	}
}
