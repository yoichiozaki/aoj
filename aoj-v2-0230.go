package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

type P struct {
	first, second int
}

var (
	n int
	t [200][2]int
	u [200][2]int
)

func change(p P) P {
	for t[p.first][p.second] == 1 && t[p.first+1][p.second] == 1 {
		p.first++
	}
	for t[p.first][p.second] == 2 {
		p.first--
	}
	return p
}

func solve() {
	Q := make([]P, 0)
	a := P{0, 0}
	b := P{0, 1}
	a = change(a)
	b = change(b)
	u[a.first][a.second] = 0
	u[b.first][b.second] = 0
	Q = append(Q, a)
	Q = append(Q, b)
	for len(Q) > 0 {
		p := Q[0]
		Q = Q[1:]
		p = change(p)
		if p.first == n-1 {
			fmt.Println(u[p.first][p.second])
			return
		}
		for i := 0; i < 3; i++ {
			np := P{p.first + i, 1 - p.second}
			np = change(np)
			if u[np.first][np.second] == -1 && np.first <= n {
				u[np.first][np.second] = u[p.first][p.second] + 1
				Q = append(Q, np)
			}
		}
	}
	fmt.Println("NA")
}

func _init() {
	for j := 0; j < 2; j++ {
		for i := 0; i < 200; i++ {
			t[i][j] = 0
			u[i][j] = -1
		}
	}
}

func main() {
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		_init()
		for j := 0; j < 2; j++ {
			for i := 0; i < n; i++ {
				scanner.Scan()
				t[i][j], _ = strconv.Atoi(scanner.Text())
			}
		}
		solve()
	}
}
