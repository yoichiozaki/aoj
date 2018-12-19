package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 0x7FFFFFFF

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	r, _ := strconv.Atoi(buf[0])
	g, _ := strconv.Atoi(buf[1])
	b, _ := strconv.Atoi(buf[2])
	total_min := INF
	cnt_r_min := INF
	cnt_b_min := INF
	var y_tmp, z_tmp int
	ans := make([]int, 3)
	for x := -300; x < 300; x++ {
		cnt_g := calc(g, x)
		for y := -r; y < 500; y++ {
			if r-y-101 >= -x {
				continue
			} else {
				cnt_r := calc(r, y)
				cnt_r_min = cnt_r
				y_tmp = y
			}
		}
		for z := -500; z < b; z++ {
			if 100-z <= g-x-1 {
				continue
			} else {
				cnt_b := calc(b, z)
				if cnt_b_min > cnt_b {
					cnt_b_min = cnt_b
					z_tmp = z
				}
			}
		}
		total := cnt_g + cnt_r_min + cnt_b_min
		if total_min > total {
			total_min = total
			ans[0] = y_tmp
			ans[1] = x
			ans[2] = z_tmp
		}
		cnt_r_min = INF
		cnt_b_min = INF
	}
	fmt.Println(total_min)
}

func calc(m, x int) int {
	// m個のマーブルを、その左端が中央から左にx番目となるように
	// 置いた時にかかるコストを計算する
	if m <= x {
		x1 := (x * (x + 1)) / 2
		x2 := ((x - m) * (x - m + 1)) / 2
		return x1 - x2
	}
	if x < 0 {
		x2 := (x * (x + 1)) / 2
		x1 := ((x - m) * (x - m + 1)) / 2
		return x1 - x2
	} else {
		x1 := (x * (x + 1)) / 2
		x2 := ((m - x - 1) * (m - x)) / 2
		return x1 + x2
	}
}
