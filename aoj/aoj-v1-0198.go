package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	// 上面の色を固定してから、左右方向に回転して一致するものがあるかを調べる
	round := [][]int{
		{1, 2, 3, 4, 5, 6}, {1, 3, 5, 2, 4, 6}, {1, 4, 2, 5, 3, 6}, {1, 5, 4, 3, 2, 6}, // c1を上面に固定して回転
		{2, 6, 3, 4, 1, 5}, {2, 3, 1, 6, 4, 5}, {2, 1, 4, 3, 6, 5}, {2, 4, 6, 1, 3, 5}, // c2を上面に固定して回転
		{3, 1, 2, 5, 6, 4}, {3, 2, 6, 1, 5, 4}, {3, 5, 1, 6, 2, 4}, {3, 6, 5, 2, 1, 4}, // c3を上面に固定して回転
		{4, 1, 5, 2, 6, 3}, {4, 2, 1, 6, 5, 3}, {4, 5, 6, 1, 2, 3}, {4, 6, 2, 5, 1, 3}, // c4を上面に固定して回転
		{5, 1, 3, 4, 6, 2}, {5, 3, 6, 1, 4, 2}, {5, 4, 1, 6, 3, 2}, {5, 6, 4, 3, 1, 2}, // c5を上面に固定して回転
		{6, 2, 4, 3, 5, 1}, {6, 3, 2, 5, 4, 1}, {6, 5, 3, 4, 2, 1}, {6, 4, 5, 2, 3, 1}, // c6を上面に固定して回転

	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		cubes := make([][]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			cubes[i] = strings.Split(scanner.Text(), " ")
		}
		check := func(a, b int) bool {
			tmp := make([]string, 6)
			for i := 0; i < 24; i++ {
				for j := 0; j < 6; j++ {
					tmp[j] = cubes[b][round[i][j]-1]
				}
				if reflect.DeepEqual(tmp, cubes[a]) {
					return true
				}
			}
			return false
		}
		tbl := make([]int, n)
		// 全ての組み合わせを調べる
		for i := 0; i < n; i++ {
			if tbl[i] == 1 {
				continue
			}
			for j := i + 1; j < n; j++ {
				if check(i, j) {
					tbl[j] = 1
				}
			}
		}
		fmt.Println(sum(tbl))
	}
}

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}
