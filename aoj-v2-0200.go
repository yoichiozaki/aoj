// Warshall–Floyd Algorithm

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 2000000000

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		m, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		c_matrix := make([][]int, m)
		for i := range c_matrix {
			c_matrix[i] = make([]int, m)
		}
		t_matrix := make([][]int, m)
		for i := range t_matrix {
			t_matrix[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(buf_[0])
			b, _ := strconv.Atoi(buf_[1])
			c, _ := strconv.Atoi(buf_[2])
			t, _ := strconv.Atoi(buf_[3])
			c_matrix[a-1][b-1] = c
			c_matrix[b-1][a-1] = c
			t_matrix[a-1][b-1] = t
			t_matrix[b-1][a-1] = t
		}
		c_dist := make([][]int, m)
		c_prev := make([][]int, m)
		t_dist := make([][]int, m)
		t_prev := make([][]int, m)
		for i := 0; i < m; i++ {
			c_dist[i] = make([]int, m)
			c_prev[i] = make([]int, m)
			t_dist[i] = make([]int, m)
			t_prev[i] = make([]int, m)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				if c_matrix[i][j] == 0 {
					c_dist[i][j] = INF
					c_prev[i][j] = -1
				} else {
					c_dist[i][j] = c_matrix[i][j]
					c_prev[i][j] = i
				}
				if t_matrix[i][j] == 0 {
					t_dist[i][j] = INF
					t_prev[i][j] = -1
				} else {
					t_dist[i][j] = t_matrix[i][j]
					t_prev[i][j] = i
				}
			}
			c_dist[i][i] = 0
			t_dist[i][i] = 0
		}
		for k := 0; k < m; k++ {
			for i := 0; i < m; i++ {
				for j := 0; j < m; j++ {
					c_newlen := c_dist[i][k] + c_dist[k][j]
					if c_newlen < c_dist[i][j] {
						c_dist[i][j] = c_newlen
						c_prev[i][j] = c_prev[k][j]
					}
				}
			}
		}
		for k := 0; k < m; k++ {
			for i := 0; i < m; i++ {
				for j := 0; j < m; j++ {
					t_newlen := t_dist[i][k] + t_dist[k][j]
					if t_newlen < t_dist[i][j] {
						t_dist[i][j] = t_newlen
						t_prev[i][j] = t_prev[k][j]
					}

				}
			}
		}
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < k; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			p, _ := strconv.Atoi(buf_[0])
			q, _ := strconv.Atoi(buf_[1])
			r, _ := strconv.Atoi(buf_[2])
			switch r {
			case 0:
				fmt.Println(c_dist[p-1][q-1])
			case 1:
				fmt.Println(t_dist[p-1][q-1])
			}
		}
	}
}
