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

var (
	R, C    int
	maze    [][]string
	cost    [][]int
	visited [][]bool
)

func main() {
	R, C = scan()
	sy, sx := scan()
	gy, gx := scan()
	sy--
	sx--
	gy--
	gx--
	maze = make([][]string, R)
	for i := range maze {
		scanner.Scan()
		maze[i] = strings.Split(scanner.Text(), "")
	}
	cost = make([][]int, R)
	for i := range cost {
		cost[i] = make([]int, C)
	}
	visited = make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}
	bfs(maze, cost, visited, sx, sy, gx, gy)
	fmt.Println(cost[gy][gx])
}

func scan() (int, int) {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	y, _ := strconv.Atoi(buf[0])
	x, _ := strconv.Atoi(buf[1])
	return y, x
}

func bfs(maze [][]string, cost [][]int, visited [][]bool, sx, sy, gx, gy int) {
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}
	if sx == gx && sy == gy {
		return
	}
	visited[sy][sx] = true
	cost[sy][sx] = 0
	que := make([]Point, 0)
	que = append(que, Point{sx, sy})
	for len(que) > 0 {
		current := que[0]
		que = que[1:]
		for i := 0; i < 4; i++ {
			next := Point{current.x + dx[i], current.y + dy[i]}
			if 0 <= next.x && next.x < C &&
				0 <= next.y && next.y < R &&
				maze[next.y][next.x] == "." && !visited[next.y][next.x] {
				cost[next.y][next.x] = cost[current.y][current.x] + 1
				visited[next.y][next.x] = true
				que = append(que, next)
			} else {
				continue
			}
		}
	}
}
