package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 0x7FFFFFFF

var pass [][]int

func main() {
	for scanner.Scan() {
		pass = make([][]int, 200)
		for i := range pass {
			pass[i] = make([]int, 200)
		}
		buf := strings.Split(scanner.Text(), " ")
		c, _ := strconv.Atoi(buf[0])
		n, _ := strconv.Atoi(buf[1])
		m, _ := strconv.Atoi(buf[2])
		s, _ := strconv.Atoi(buf[3])
		d, _ := strconv.Atoi(buf[4])
		if c == 0 && n == 0 && m == 0 && s == 0 && d == 0 {
			break
		}
		s--
		d--
		for i := 0; i < 200; i++ {
			for j := 0; j < 200; j++ {
				pass[i][j] = INF
			}
		}
		for i := 0; i < 200; i++ {
			pass[i][i] = 0
		}
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(buf_[0])
			b, _ := strconv.Atoi(buf_[1])
			f, _ := strconv.Atoi(buf_[2])
			a--
			b--
			pass[a][b] = f
			pass[b][a] = f
		}
		fmt.Println(dijkstra(s, d, c, n))
	}
}

func dijkstra(s, g, c, n int) float64 {
	d := make([][]float64, 200)
	used := make([][]bool, 200)
	for i := 0; i < 200; i++ {
		d[i] = make([]float64, 20)
		used[i] = make([]bool, 20)
	}
	for i := range d {
		for j := range d[i] {
			d[i][j] = INF
		}
	}
	d[s][c] = 0
	for {
		v := -1
		w := -1
		for u := 0; u < n; u++ {
			for i := 0; i < c+1; i++ {
				if !used[u][i] && (v == -1 || d[u][i] < d[v][w]) {
					v = u
					w = i
				}
			}
		}
		if v == -1 {
			break
		}
		used[v][w] = true
		for u := 0; u < n; u++ {
			d[u][w] = math.Min(d[u][w], d[v][w]+float64(pass[v][u]))
			if w-1 >= 0 && pass[v][u] != INF {
				d[u][w-1] = math.Min(d[u][w-1], d[v][w]+(float64(pass[v][u]/2.0)))
			}
		}
	}
	ret := float64(INF)
	for i := 0; i < c+1; i++ {
		ret = math.Min(ret, d[g][i])
	}
	return ret
}
