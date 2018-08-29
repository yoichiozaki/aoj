// BFS

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var p = [15]int{}
var g = [13][4]int{
	// -1は存在しないマス目
	{2, -1, -1, -1},
	{2, 5, -1, -1}, {0, 1, 3, 6}, {2, 7, -1, -1},
	{5, -1, -1, -1}, {1, 4, 6, 9}, {2, 5, 7, 10}, {3, 6, 8, 11}, {7, -1, -1, -1},
	{5, 10, -1, -1}, {6, 9, 11, 12}, {7, 10, -1, -1},
	{10, -1, -1, -1},
}

/*
　  ０ １ ２ ３ ４
 0 |／|／|  |／|／|
 1 |／|　|  |　|／|
 2 |　|　|  |　|　|
 3 |／|　|  |　|／|
 4 |／|／|  |／|／|
*/

var row = [13]int{0, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 4}
var col = [13]int{2, 1, 2, 3, 0, 1, 2, 3, 4, 1, 2, 3, 2}

type wolf struct {
	m [13]int
}

type pair struct {
	first  wolf
	second int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func print_board(board [13]int) {
	fmt.Println("--------------------------")
	fmt.Printf("        [%2d]\n", board[0])
	fmt.Printf("    [%2d][%2d][%2d]\n", board[1], board[2], board[3])
	fmt.Printf("[%2d][%2d][%2d][%2d][%2d]\n", board[4], board[5], board[6], board[7], board[8])
	fmt.Printf("    [%2d][%2d][%2d]\n", board[9], board[10], board[11])
	fmt.Printf("        [%2d]\n", board[12])
	fmt.Println("--------------------------")
}

func main() {
	for scanner.Scan() {
		p1, _ := strconv.Atoi(scanner.Text())
		if p1 == -1 {
			break
		}
		p[0] = p1
		scanner.Scan()
		buf1 := strings.Split(scanner.Text(), " ")
		p2, _ := strconv.Atoi(buf1[0])
		p3, _ := strconv.Atoi(buf1[1])
		p4, _ := strconv.Atoi(buf1[2])
		p[1], p[2], p[3] = p2, p3, p4
		scanner.Scan()
		buf2 := strings.Split(scanner.Text(), " ")
		p5, _ := strconv.Atoi(buf2[0])
		p6, _ := strconv.Atoi(buf2[1])
		p7, _ := strconv.Atoi(buf2[2])
		p8, _ := strconv.Atoi(buf2[3])
		p9, _ := strconv.Atoi(buf2[4])
		p[4], p[5], p[6], p[7], p[8] = p5, p6, p7, p8, p9
		scanner.Scan()
		buf3 := strings.Split(scanner.Text(), " ")
		p10, _ := strconv.Atoi(buf3[0])
		p11, _ := strconv.Atoi(buf3[1])
		p12, _ := strconv.Atoi(buf3[2])
		p[9], p[10], p[11] = p10, p11, p12
		scanner.Scan()
		p13, _ := strconv.Atoi(scanner.Text())
		p[12] = p13
		S := make(map[wolf]struct{})
		st := wolf{}
		for i := 0; i < 13; i++ {
			st.m[i] = p[i]
		}
		Q := make([]pair, 0)
		S[st] = struct{}{}
		Q = append(Q, pair{st, 0})
		yet := true
		for len(Q) > 0 {
			at := Q[0].first
			// print_board(at.m)
			cost := Q[0].second
			Q = Q[1:]
			ok := true
			for i := 1; i < 12; i++ {
				if at.m[i] != i {
					ok = false
				}
			}
			if ok {
				// print_board(at.m)
				yet = false
				fmt.Println(cost)
				break
			}
			now := cost
			for i := 0; i < 13; i++ {
				if at.m[i] != 0 {
					now += abs(row[i]-row[at.m[i]]) + abs(col[i]-col[at.m[i]])
				}
			}
			if now > 20 {
				continue
			}
			for i := 0; i < 13; i++ {
				if at.m[i] == 0 {
					for j := 0; j < 4; j++ {
						if ^g[i][j] != 0 {
							at.m[i], at.m[g[i][j]] = at.m[g[i][j]], at.m[i]
							if _, ok := S[at]; !ok {
								S[at] = struct{}{}
								Q = append(Q, pair{at, cost + 1})
							}
							at.m[i], at.m[g[i][j]] = at.m[g[i][j]], at.m[i]
						}
					}
				}
			}
		}
		if yet {
			fmt.Println("NA")
		}
	}
}
