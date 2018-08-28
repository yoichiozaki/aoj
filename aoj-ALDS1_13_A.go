
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const (
	N = 8
	FREE = -1
	ATTACKED = 1
)

var scanner = bufio.NewScanner(os.Stdin)

var (
	row []int // row[x](x = 0, 1, ... ,N-1)がATTACKEDなら行xは襲撃されている
	col []int // col[x](x = 0, 1, ... ,N-1)がATTACKEDなら列xは襲撃されている
	dpos []int // dpos[x](x = 0, 1, ... ,2(N-1))がATTACKEDなら左下方向の列xは襲撃されている
	dneg []int // dneg[x](x = 0, 1, ... ,2(N-1))がATTACKEDなら右下方向の列xは襲撃されている
	X [][]bool // 目当ての盤目を指定するための二次配列
)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func initialize() {
	row = make([]int, N)
	col = make([]int ,N)
	dpos = make([]int, 2*N-1)
	dneg = make([]int, 2*N-1)
	for i := 0; i < N; i++ {
		row[i] = FREE
		col[i] = FREE
	}
	for i := 0; i < 2*N-1; i++ {
		dpos[i] = FREE
		dneg[i] = FREE
	}
}

func printBoard() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if X[i][j] {
				if row[i] != j {
					return
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if row[i] == j {
				fmt.Printf("Q")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func solve(i int) {
	if i == N {
		printBoard()
		return
	}
	for j := 0; j < N; j++ {
		if ATTACKED == col[j] || ATTACKED == dpos[i+j] || ATTACKED == dneg[i-j+N-1] {
			continue
		}
		row[i] = j
		col[j] = ATTACKED
		dpos[i+j] = ATTACKED
		dneg[i-j+N-1] = ATTACKED
		solve(i+1)
		row[i] = FREE
		col[j] = FREE
		dpos[i+j] = FREE
		dneg[i-j+N-1] = FREE
	}
}

func main() {
	scanner.Split(bufio.ScanWords)
	k := nextInt()
	initialize()
	X = make([][]bool, N)
	for i := 0; i < N; i++ {
		X[i] = make([]bool, N)
		for j := 0; j < N; j++ {
			X[i][j] = false
		}
	}
	for i := 0; i < k; i++ {
		r := nextInt()
		c := nextInt()
		X[r][c] = true
	}
	solve(0)

}
