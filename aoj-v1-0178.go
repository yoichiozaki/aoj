package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var (
	dat = [5000][5]int{}
	m   = [5]int{}
)

func main() {
	for scanner.Scan() {
		for i := range dat {
			for j := range dat[i] {
				dat[i][j] = 0
			}
		}
		for i := range m {
			m[i] = 0
		}
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		count := 0
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			d, _ := strconv.Atoi(buf[0])
			p, _ := strconv.Atoi(buf[1])
			q, _ := strconv.Atoi(buf[2])
			q--
			count += p
			if d == 2 {
				for j := 0; j < p; j++ {
					dat[m[q]][q] = 1
					ok := true
					for k := 0; k < 5; k++ {
						if dat[m[q]][k] == 0 {
							ok = false
						}
					}
					if ok {
						for k := 0; k < 5; k++ {
							dat[m[q]][k] = 0
							for l := m[q]; l < 4999; l++ {
								dat[l][k] = dat[l+1][k]
							}
							for l := 4999; l >= 0; l-- {
								if dat[l][k] != 0 {
									m[k] = l + 1
									break
								}
								if l == 0 {
									m[k] = 0
								}
							}
						}
						count -= 5
					} else {
						m[q]++
					}
				}
			} else {
				h := 0
				for i := 0; i < p; i++ {
					h = max(h, m[i+q])
				}
				for i := 0; i < p; i++ {
					dat[h][i+q] = 1
				}
				for i := 0; i < p; i++ {
					m[i+q] = h + 1
				}
				ok := true
				for j := 0; j < 5; j++ {
					if dat[h][j] == 0 {
						ok = false
					}
				}
				if ok {
					for j := 0; j < 5; j++ {
						dat[h][j] = 0
						for k := h; k < 4999; k++ {
							dat[k][j] = dat[k+1][j]
						}
						for k := 4999; k >= 0; k-- {
							if dat[k][j] != 0 {
								m[j] = k + 1
								break
							}
							if k == 0 {
								m[j] = 0
							}
						}
					}
					count -= 5
				}
			}
		}
		fmt.Println(count)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
