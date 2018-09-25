package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type pair struct {
	first, second int
}

var (
	g             [20][]pair
	ret, now, dim int
)

func main() {
	for {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		ret, now, dim = 0, 0, 0
		for i := 0; i < 20; i++ {
			g[i] = make([]pair, 0)
		}
		for i := 0; i < a-1; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			b, _ := strconv.Atoi(buf[0])
			c, _ := strconv.Atoi(buf[1])
			d, _ := strconv.Atoi(buf[2])
			b--
			c--
			g[b] = append(g[b], pair{c, d})
			g[c] = append(g[c], pair{b, d})
		}
		dfs(0, -1, 0)
		fmt.Println(ret - dim)
	}
}

func dfs(a, b, c int) {
	if a != 0 && len(g[a]) == 1 {
		return
	}
	ret += c * 2
	now += c
	dim = max(dim, now)
	for i := 0; i < len(g[a]); i++ {
		if g[a][i].first != b {
			dfs(g[a][i].first, a, g[a][i].second)
		}
	}
	now -= c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
