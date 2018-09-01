package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var N, M, res int
var member []int
var check []bool
var G [][]int

func dfs(v, c int) {
	log.Println(member)
	check[v] = true
	c++
	for _, e := range G[v] {
		member[e]++
	}
	if c > res {
		res = c
	}

	for _, next := range G[v] {
		if check[next] == false && member[next] == c {
			dfs(next, c)
		}
	}

	for _, e := range G[v] {
		member[e]--
	}
}

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")

	N, _ := strconv.Atoi(buf[0])
	M, _ := strconv.Atoi(buf[1])

	G = make([][]int, N+1)
	for i := 0; i < M; i++ {
		scanner.Scan()
		buf_ := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(buf_[0])
		y, _ := strconv.Atoi(buf_[1])
		G[x] = append(G[x], y) // xとyは知り合い
		G[y] = append(G[y], x) // yとxは知り合い
	}
	log.Println(G)

	for i := 1; i <= N; i++ {
		member = make([]int, N+1)
		check = make([]bool, N+1)
		dfs(i, 0)
	}

	fmt.Println(res)
}
