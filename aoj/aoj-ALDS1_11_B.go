
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/***************************************************************
stackによる深さ優先探索を実装する。
***************************************************************/

const (
	N     = 100
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

var (
	n            int
	Adj          [N][N]int
	statusColor  [N]int
	discoverTime [N]int
	finishTime   [N]int
	clock        int
	nt           [N]int
)

/***************************************************************
Stackの定義:
- field
	container: ノードの格納されたスライス
	sp: スタックポインタ
- method
	size(): スタックに積まれているノードの個数を返す
	top(): スタックの一番上に積まれているノードを返す
	isEmpty(): スタックが空かどうかを返す
	pop(): スタックの要素をポップし削除する
	push(): スタックに要素を積む
***************************************************************/

type Stack struct {
	container []int
	sp        int
}

func (s *Stack) size() int {
	return len(s.container)
}

func (s *Stack) top() int {
	return s.container[s.sp]
}

func (s *Stack) isEmpty() bool {
	return s.sp == 0
}

func (s *Stack) pop() int {
	_pop := func() {
		s.container = s.container[:len(s.container)-1]
	}
	defer _pop()
	if s.isEmpty() {
		panic("There is no elements in stack!")
	}
	s.sp--
	return s.container[s.sp]
}

func (s *Stack) push(n int) {
	s.sp++
	s.container[s.sp] = n
}

func next(u int) int {
	for v := nt[u]; v < n; v++ {
		nt[u] = v + 1
		if Adj[u][v] == 1 {
			return v
		}
	}
	return -1
}

func _DFS(target int) {
	for i := 0; i < n; i++ {
		nt[i] = 0
	}
	stack := new(Stack)
	stack.container = make([]int, N+1)
	stack.push(target)
	statusColor[target] = GRAY
	clock++
	discoverTime[target] = clock
	for !stack.isEmpty() {
		u := stack.top()
		v := next(u)
		if v != -1 {
			if statusColor[v] == WHITE {
				statusColor[v] = GRAY
				clock++
				discoverTime[v] = clock
				stack.push(v)
			}
		} else {
			stack.pop()
			statusColor[u] = BLACK
			clock++
			finishTime[u] = clock
		}
	}
}

func DFS() {
	for i := 0; i < n; i++ {
		statusColor[i] = WHITE
		nt[i] = 0
	}
	clock = 0
	for i := 0; i < n; i++ {
		if statusColor[i] == WHITE {
			_DFS(i)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d %d %d\n", i+1, discoverTime[i], finishTime[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Adj[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		u--
		for j := 0; j < k; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			v--
			Adj[u][v] = 1
		}
	}

	DFS()
}

