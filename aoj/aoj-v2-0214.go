package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

var (
	uf_table []int
	x        [][]int
	y        [][]int
)

func main() {
	scanner.Split(bufio.ScanWords)
	uf_table = make([]int, 100)
	x = make([][]int, 100)
	for i := range x {
		x[i] = make([]int, 4)
	}
	y = make([][]int, 100)
	for i := range y {
		y[i] = make([]int, 4)
	}
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < N; i++ {
		scanner.Scan()
		M, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < M; j++ {
			for k := 0; k < 4; k++ {
				scanner.Scan()
				x[j][k], _ = strconv.Atoi(scanner.Text())
				scanner.Scan()
				y[j][k], _ = strconv.Atoi(scanner.Text())
			}
		}
		for j := 0; j < M; j++ {
			uf_table[j] = -1
		}
		for j := 0; j < M; j++ {
			for k := j + 1; k < M; k++ {
				ok := false
				for l := 0; l < 4; l++ {
					if inclusive(x[j][0], y[j][0], x[j][1], y[j][1], x[j][2], y[j][2], x[k][l], y[k][l]) {
						ok = true
					}
					if inclusive(x[j][0], y[j][0], x[j][3], y[j][3], x[j][2], y[j][2], x[k][l], y[k][l]) {
						ok = true
					}
					if inclusive(x[k][0], y[k][0], x[k][1], y[k][1], x[k][2], y[k][2], x[j][l], y[j][l]) {
						ok = true
					}
					if inclusive(x[k][0], y[k][0], x[k][3], y[k][3], x[k][2], y[k][2], x[j][l], y[j][l]) {
						ok = true
					}
				}
				for l := 0; l < 4; l++ {
					for m := 0; m < 4; m++ {
						if intersect(x[j][l], y[j][l], x[j][(l+1)%4], y[j][(l+1)%4], x[k][m], y[k][m], x[k][(m+1)%4], y[k][(m+1)%4]) {
							ok = true
						}
					}
				}
				if ok {
					union(j, k)
				}
			}
		}
		ret := 0
		for j := 0; j < M; j++ {
			if uf_table[i] < 0 {
				ret++
			}
		}
		fmt.Println(ret)
	}
}

func find(a int) int {
	if uf_table[a] < 0 {
		return a
	}
	uf_table[a] = find(uf_table[a])
	return uf_table[a]
}

func union(a, b int) {
	a = find(a)
	b = find(b)
	if a == b {
		return
	}
	uf_table[a] += uf_table[b]
	uf_table[b] = a
}

func inclusive(ax, ay, bx, by, cx, cy, px, py int) bool {
	ok1 := true
	ok2 := true
	if (bx-ax)*(py-ay)-(px-ax)*(by-ay) < 0 {
		ok1 = false
	}
	if (bx-ax)*(py-ay)-(px-ax)*(by-ay) > 0 {
		ok2 = false
	}
	if (cx-bx)*(py-by)-(px-bx)*(cy-by) < 0 {
		ok1 = false
	}
	if (cx-bx)*(py-by)-(px-bx)*(cy-by) > 0 {
		ok2 = false
	}
	if (ax-cx)*(py-cy)-(px-cx)*(ay-cy) < 0 {
		ok1 = false
	}
	if (ax-cx)*(py-cy)-(px-cx)*(ay-cy) > 0 {
		ok2 = false
	}
	return ok1 || ok2
}

func intersect(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	ta := (cx-dx)*(ay-cy) + (cy-dy)*(cx-ax)
	tb := (cx-dx)*(by-cy) + (cy-dy)*(cx-bx)
	tc := (ax-bx)*(cy-ay) + (ay-by)*(ax-cx)
	td := (ax-bx)*(dy-ay) + (ay-by)*(ax-dx)
	if ta*tb < 0 && tc*td < 0 {
		return true
	}
	return false
}
