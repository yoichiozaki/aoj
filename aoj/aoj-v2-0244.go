package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 1 << 29

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		m, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		d := make([][]int, 100)
		for i := range d {
			d[i] = make([]int, 100)
		}
		b := make([][]bool, 100)
		for i := range d {
			b[i] = make([]bool, 100)
		}
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if u == v {
					d[u][v] = 0
				} else {
					d[u][v] = INF
				}
			}
		}
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			u, _ := strconv.Atoi(buf_[0])
			v, _ := strconv.Atoi(buf_[1])
			c, _ := strconv.Atoi(buf_[2])
			u--
			v--
			d[u][v], d[v][u] = min(d[u][v], c), min(d[u][v], c)
			b[u][v], b[v][u] = true, true
		}
		for k := 0; k < n; k++ {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					d[i][j] = min(d[i][j], d[i][k]+d[k][j])
				}
			}
		}
		ans := d[0][n-1]
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				ok := false
				for w := 0; w < n; w++ {
					if w != u && w != v && b[u][w] && b[w][v] {
						ok = true
						break
					}
				}
				if ok {
					ans = min(ans, d[0][u]+d[v][n-1])
				}
			}
		}
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
