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
	dx = []int{0, -1, 0, 1}
	dy = []int{1, 0, -1, 0}
)

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		w, _ := strconv.Atoi(buf[0])
		h, _ := strconv.Atoi(buf[1])
		if w == 0 && h == 0 {
			break
		}
		scanner.Scan()
		buf_ := strings.Split(scanner.Text(), " ")
		xs, _ := strconv.Atoi(buf_[0])
		ys, _ := strconv.Atoi(buf_[1])
		scanner.Scan()
		buf__ := strings.Split(scanner.Text(), " ")
		xg, _ := strconv.Atoi(buf__[0])
		yg, _ := strconv.Atoi(buf__[1])
		xs--
		ys--
		xg--
		yg--
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		field := make([][]int, 100)
		for i := range field {
			field[i] = make([]int, 100)
		}
		visited := make([][]bool, 100)
		for i := range visited {
			visited[i] = make([]bool, 100)
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf___ := strings.Split(scanner.Text(), " ")
			c, _ := strconv.Atoi(buf___[0])
			d, _ := strconv.Atoi(buf___[1])
			x, _ := strconv.Atoi(buf___[2])
			y, _ := strconv.Atoi(buf___[3])
			x--
			y--
			switch d {
			case 0:
				for j := 0; j < 4; j++ {
					for k := 0; k < 2; k++ {
						field[x+j][y+k] = c
					}
				}
			case 1:
				for j := 0; j < 2; j++ {
					for k := 0; k < 4; k++ {
						field[x+j][y+k] = c
					}
				}
			}
		}
		if field[xs][ys] != field[xg][yg] || // スタートとゴールの色が等しくないと"NG"
			field[xs][ys] == 0 { // スタートの色が無職
			fmt.Println("NG")
			continue
		}
		Q := make([]Entry, 0)
		visited[xs][ys] = true
		Q = append(Q, Entry{xs, ys})
		for len(Q) > 0 {
			row := Q[0].x
			col := Q[0].y
			Q = Q[1:]
			for i := 0; i < 4; i++ {
				if 0 <= row+dx[i] && row+dx[i] < w &&
					0 <= col+dy[i] && col+dy[i] < h && // 移動可能な範囲内かどうか
					!visited[row+dx[i]][col+dy[i]] && // 訪問済みか
					field[row+dx[i]][col+dy[i]] == field[row][col] { // 色が同じかどうか
					visited[row+dx[i]][col+dy[i]] = true
					Q = append(Q, Entry{row + dx[i], col + dy[i]})
				}
			}
		}
		if visited[xg][yg] {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}

type Entry struct {
	x, y int
}
