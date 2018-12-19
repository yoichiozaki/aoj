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
	h, w    int
	maze    [][]string
	visited [][]bool
)

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	h, _ = strconv.Atoi(buf[0])
	w, _ = strconv.Atoi(buf[1])
	maze = make([][]string, h)
	for i := range maze {
		scanner.Scan()
		maze[i] = strings.Split(scanner.Text(), "")
	}
	visited = make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	sx, sy, gx, gy := 0, 0, 0, 0
	for i := range maze {
		for j := range maze[i] {
			switch maze[i][j] {
			case "s":
				sx, sy = j, i
			case "g":
				gx, gy = j, i
			}
		}
	}
	search(sx, sy)
	if visited[gy][gx] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func search(x, y int) {
	if x < 0 || w <= x || y < 0 || h <= y || maze[y][x] == "#" {
		return
	}
	if visited[y][x] {
		return
	}
	visited[y][x] = true
	search(x+1, y)
	search(x-1, y)
	search(x, y+1)
	search(x, y-1)
}
