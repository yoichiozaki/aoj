package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Point struct {
	x, y int
}

func main() {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	maze := make([][]rune, 30)
	for i := range maze {
		maze[i] = make([]rune, 30)
	}
	next := make([][]rune, 30)
	for i := range next {
		next[i] = make([]rune, 30)
	}
	check := make([][]int, 30)
	for i := range check {
		check[i] = make([]int, 30)
	}
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		W, _ := strconv.Atoi(buf[0])
		H, _ := strconv.Atoi(buf[1])
		if W == 0 && H == 0 {
			break
		}
		for i := 0; i < H; i++ {
			scanner.Scan()
			buf_ := scanner.Text()
			for j, v := range buf_ {
				maze[i][j] = v
			}
		}
		man := 0
		for i := 0; i < H; i++ {
			for j := 0; j < H; j++ {
				if maze[i][j] != '.' && maze[i][j] != 'X' && maze[i][j] != '#' {
					switch maze[i][j] {
					case 'E':
						maze[i][j] = 0
						man++
					case 'N':
						maze[i][j] = 1
						man++
					case 'W':
						maze[i][j] = 2
						man++
					case 'S':
						maze[i][j] = 3
						man++
					}
				}
			}
		}
		ans := -1
		for t := 0; t <= 180; t++ {
			if man == 0 {
				ans = t
				break
			}
			for i := 0; i < H; i++ {
				for j := 0; j < W; j++ {
					check[i][j] = 0x7FFFFFFF
				}
			}
			for i := 0; i < H; i++ {
				for j := 0; j < W; j++ {
					if maze[i][j] > 3 {
						continue
					}
					l := (maze[i][j] + 3) % 4
					for k := 0; k < 4; k++ {
						y := i + dy[l]
						x := j + dx[l]
						if 0 <= y && y < H && 0 <= x && x < W && (maze[y][x] == '.' || maze[y][x] == 'X') {
							maze[i][j] = l
							check[y][x] = min(check[y][x], int((l+2)%4))
							break
						}
						l = (l + 1) % 4
					}
				}
			}
			for i := 0; i < H; i++ {
				for j := 0; j < W; j++ {
					next[i][j] = maze[i][j]
				}
			}
			for i := 0; i < H; i++ {
				for j := 0; j < W; j++ {
					if maze[i][j] > 3 {
						continue
					}
					y := i + dy[maze[i][j]]
					x := j + dx[maze[i][j]]
					if 0 <= y && y < H && 0 <= x && x < W && check[y][x] == int((maze[i][j]+2)%4) {
						if maze[y][x] == '.' {
							next[y][x] = maze[i][j]
						} else if maze[y][x] == 'X' {
							man--
						}
						next[i][j] = '.'
					}
				}
			}
			for i := 0; i < H; i++ {
				for j := 0; j < W; j++ {
					maze[i][j] = next[i][j]
				}
			}
		}
		if ^ans != 0 {
			fmt.Println(ans)
		} else {
			fmt.Println("NA")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
