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
	n, _ := strconv.Atoi(buf[0])
	m, _ := strconv.Atoi(buf[1])
	e := make([][]int, n)
	for i := range e {
		e[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		buf_ := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(buf_[0])
		y, _ := strconv.Atoi(buf_[1])
		x--
		y--
		e[x][y], e[y][x] = 1, 1
	}
	ans := 0
	for bit := 1; bit < (1 << uint(n)); bit++ { // 2^n通り全て調べる
		t := bitCount(bit)
		if t <= ans { // 全部0はあり得ないし、最大のものを求めたい
			continue
		}
		flag := true
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				if ((uint(bit)>>uint(i))&(uint(bit)>>uint(j))&1) != 0 &&
					e[i][j] != 1 {
					flag = false
				}
			}
		}
		if flag {
			ans = t
		}
	}
	fmt.Println(ans)
}

func bitCount(x int) int {
	cnt := 0
	strx := strings.Split(fmt.Sprintf("%012b", x), "")
	for i := range strx {
		if strx[i] == "1" {
			cnt++
		}
	}
	return cnt
}
