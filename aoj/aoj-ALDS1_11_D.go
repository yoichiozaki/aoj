
import (
	"bufio"
	"container/list"
	"os"
	"strconv"
	"fmt"
)

const (
	MAX = 100000
	NIL = -1
)

var (
	n     int
	G     [MAX][]int
	color [MAX]int
)

func DFS(r int, c int) {
	S := list.New()
	S.PushFront(r)
	color[r] = c
	for S.Len() != 0 {
		u := S.Remove(S.Front()).(int)
		for i := 0; i < len(G[u]); i++ {
			v := G[u][i]
			if color[v] == NIL {
				color[v] = c
				S.PushFront(v)
			}
		}
	}
}

func assignColor() {
	id := 1
	for i := 0; i < n; i++ {
		color[i] = NIL
	}
	for i := 0; i < n; i++ {
		if color[i] == NIL {
			id++
			DFS(i, id)
		}
	}
}

func main() {
	var s, t, m, q int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < m; i++ {
		scanner.Scan()
		s, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		t, _ = strconv.Atoi(scanner.Text())
		if i == 0 {
			G[s] = make([]int, 0)
			G[t] = make([]int, 0)
		}
		G[s] = append(G[s], t)
		G[t] = append(G[t], s)
	}

	assignColor()
	scanner.Scan()
	q, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < q; i++ {
		scanner.Scan()
		s, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		t, _ = strconv.Atoi(scanner.Text())
		if color[s] == color[t] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

