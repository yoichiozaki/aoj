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

var (
	direction1 = [][]int{ // 奇数段目の移動方向
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, -1}, {1, 1},
	}
	direction2 = [][]int{ // 偶数段目の移動方向
		{-1, -1}, {0, -1}, {1, 0}, {0, 1}, {-1, 1}, {-1, 0},
	}
)

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(buf[0])
		n, _ := strconv.Atoi(buf[1])
		if m == 0 && n == 0 {
			break
		}
		// initialize district_map
		/*
			district_map = [
				[-1, -1,  -1,  -1,  -1]
				[-1, INF, INF, INF, -1]
				[-1, INF, INF, INF, -1]
				[-1, INF, INF, INF, -1]
				[-1, INF, INF, INF, -1]
				[-1, -1,  -1,  -1,  -1]
			]
		*/
		district_map := make([][]int, n+2)
		for i := range district_map {
			district_map[i] = make([]int, m+2)
		}
		tmp1 := make([]int, m+2)
		for i := range tmp1 {
			tmp1[i] = -1
		}
		district_map[0] = tmp1
		district_map[n+1] = tmp1
		for i := 1; i < n+1; i++ {
			district_map[i] = make([]int, m+2)
			district_map[i][0] = -1
			district_map[i][m+1] = -1
			for j := 1; j < m+1; j++ {
				district_map[i][j] = INF
			}
		}

		scanner.Scan()
		s, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < s; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			x, _ := strconv.Atoi(buf_[0])
			y, _ := strconv.Atoi(buf_[1])
			Search(Point{x, y}, district_map, m, n)
		}
		ans := 0
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < t; i++ {
			scanner.Scan()
			buf__ := strings.Split(scanner.Text(), " ")
			x, _ := strconv.Atoi(buf__[0])
			y, _ := strconv.Atoi(buf__[1])
			z := Count(Point{x, y}, district_map, m, n)
			ans = max(ans, z)
		}
		fmt.Println(ans)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Point struct {
	x, y int
}

type Entry struct {
	hop   int
	point Point
}

func Search(p Point, mp [][]int, m, n int) {
	que := make([]Entry, 0)
	que = append(que, Entry{0, p})
	check := make([][]bool, n+2)
	for i := range check {
		check[i] = make([]bool, m+2)
	}
	check[p.y][p.x] = true
	mp[p.y][p.x] = 0
	var direction [][]int
	for len(que) > 0 {
		tmp := que[0]
		que = que[1:]
		dist, point := tmp.hop, tmp.point
		if point.y%2 == 0 {
			direction = direction1
		} else {
			direction = direction2
		}
		for _, d := range direction {
			nx, ny := point.x+d[0], point.y+d[1]
			if !check[ny][nx] && mp[ny][nx] != -1 {
				check[ny][nx] = true
				if mp[ny][nx] > dist+1 {
					mp[ny][nx] = dist + 1
				}
				que = append(que, Entry{dist + 1, Point{nx, ny}})
			}
		}
	}
}

func Count(p Point, mp [][]int, m, n int) int {
	que := make([]Entry, 0)
	que = append(que, Entry{0, p})
	check := make([][]bool, n+2)
	for i := range check {
		check[i] = make([]bool, m+2)
	}
	check[p.y][p.x] = true
	ans := 1
	var direction [][]int
	for len(que) > 0 {
		tmp := que[0]
		que = que[1:]
		dist, point := tmp.hop, tmp.point
		if point.y%2 == 0 {
			direction = direction1
		} else {
			direction = direction2
		}
		for _, d := range direction {
			nx, ny := point.x+d[0], point.y+d[1]
			if !check[ny][nx] && mp[ny][nx] != -1 {
				check[ny][nx] = true
				if mp[ny][nx] > dist+1 {
					ans++
				}
				que = append(que, Entry{dist + 1, Point{nx, ny}})
			}
		}
	}
	return ans
}
