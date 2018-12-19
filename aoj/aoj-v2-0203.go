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
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		X, _ := strconv.Atoi(buf[0])
		Y, _ := strconv.Atoi(buf[1])
		if X == 0 && Y == 0 {
			break
		}
		tbl := make([][]int, Y+2)
		field := make([][]int, Y+2)
		for i := range field {
			field[i] = make([]int, X+2)
			tbl[i] = make([]int, X+2)
			for j := range field[i] {
				field[i][j] = 2
				tbl[i][j] = 0
			}
		}
		buf_ := make([][]string, Y)
		for i := range buf_ {
			scanner.Scan()
			buf_[i] = strings.Split(scanner.Text(), " ")
		}
		for i := 1; i <= Y; i++ {
			for j := 1; j <= X; j++ {
				field[i][j], _ = strconv.Atoi(buf_[i-1][j-1])
				switch field[i][j] {
				case 0:
					if i == 1 {
						tbl[i][j] = 1 // スタートy=1
					} else {
						if field[i-1][j-1] != 2 {
							tbl[i][j] += tbl[i-1][j-1] // 左上からやってくる
						}
						if field[i-1][j] != 2 {
							tbl[i][j] += tbl[i-1][j] // 真上からやってくる
						}
						if field[i-1][j+1] != 2 {
							tbl[i][j] += tbl[i-1][j+1] // 右上からやってくる
						}
						if field[i-2][j] == 2 {
							tbl[i][j] += tbl[i-2][j] // ジャンプしてくる
						}
					}
				case 1:
					tbl[i][j] = 0 // 障害物あり
				case 2:
					if field[i-1][j] != 2 {
						tbl[i][j] += tbl[i-1][j] // 真上からくる
					}
					if field[i-2][j] == 2 {
						tbl[i][j] += tbl[i-2][j] // ジャンプ台からジャンプしてくる
					}
				}
			}
		}
		sum := 0
		for i := 1; i <= X; i++ {
			sum += tbl[Y][i] // とびこさないでゴール
			if field[Y-1][i] == 2 {
				sum += tbl[Y-1][i] // 飛び越してゴール
			}
		}
		fmt.Println(sum)
	}
}
