package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type State struct {
	ay, ax, by, bx, cnt int
}

var (
	H, W    int
	dy      = []int{-1, 0, 1, 0}
	dx      = []int{0, 1, 0, -1}
	t       [][]int
	visited [][][][]bool
)

func isIn(y, x int) bool {
	if y < 0 || H <= y {
		return false
	}
	if x < 0 || W <= x {
		return false
	}
	if t[y][x] == 1 {
		return false
	}
	return true
}

func solve(s State) {
	Q := make([]State, 0)
	Q = append(Q, s)
	visited[s.ay][s.ax][s.by][s.bx] = true
	for len(Q) > 0 {
		now := Q[0]
		Q = Q[1:]
		ay := now.ay
		ax := now.ax
		by := now.by
		bx := now.bx
		if now.cnt > 100 {
			continue
		}
		if ay == by && ax == bx {
			fmt.Println(now.cnt)
			return
		}
		for i := 0; i < 4; i++ {
			nay := ay + dy[i]
			nax := ax + dx[i]
			nby := by - dy[i]
			nbx := bx - dx[i]
			if !isIn(nay, nax) {
				nay = ay
				nax = ax
			}
			if !isIn(nby, nbx) {
				nby = by
				nbx = bx
			}
			if visited[nay][nax][nby][nbx] {
				continue
			}
			visited[nay][nax][nby][nbx] = true
			Q = append(Q, State{nay, nax, nby, nbx, now.cnt + 1})
		}
	}
	fmt.Println("NA")
}

func main() {
	for {
		scanner.Scan()
		buf := strings.Split(scanner.Text(), " ")
		W, _ = strconv.Atoi(buf[0])
		H, _ = strconv.Atoi(buf[1])
		if W == 0 && H == 0 {
			break
		}

		t = make([][]int, H)
		for i := range t {
			t[i] = make([]int, W)
		}

		visited = make([][][][]bool, H)
		for i := range visited {
			visited[i] = make([][][]bool, W)
			for j := range visited[i] {
				visited[i][j] = make([][]bool, H)
				for k := range visited[i][j] {
					visited[i][j][k] = make([]bool, W)
				}
			}
		}

		var s State
		scanner.Scan()
		buf = strings.Split(scanner.Text(), " ")
		s.ax, _ = strconv.Atoi(buf[0])
		s.ay, _ = strconv.Atoi(buf[1])
		scanner.Scan()
		buf = strings.Split(scanner.Text(), " ")
		s.bx, _ = strconv.Atoi(buf[0])
		s.by, _ = strconv.Atoi(buf[1])
		s.ay--
		s.ax--
		s.by--
		s.bx--
		s.cnt = 0

		for i := 0; i < H; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			for j := range buf_ {
				t[i][j], _ = strconv.Atoi(buf_[j])
			}
		}
		solve(s)
	}
}
