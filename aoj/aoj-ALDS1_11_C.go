
import (
	"container/list"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const (
	N   = 100
	INF = 1 << 21
)

var (
	n        int
	adj      [N][N]int
	distance [N]int
)

func BFS(start int) {
	q := list.New()
	q.PushBack(start)
	for i := 0; i < n; i++ {
		distance[i] = INF
	}
	distance[start] = 0
	u := 0
	for q.Len() != 0 {
		u = q.Remove(q.Front()).(int)
		for v := 0; v < n; v++ {
			if adj[u][v] == 0 {
				continue
			}
			if distance[v] != INF {
				continue
			}
			distance[v] = distance[u] + 1
			q.PushBack(v)
		}
	}
	for i := 0; i < n; i++ {
		var d int
		if distance[i] == INF {
			d = -1
		} else {
			d = distance[i]
		}
		fmt.Printf("%d %d\n", i+1, d)
	}
}

func main() {
	var u, k, v int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			adj[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		u, _ = strconv.Atoi(scanner.Text())
		u -= 1
		scanner.Scan()
		k, _ = strconv.Atoi(scanner.Text())
		for j := 0; j < k; j++ {
			scanner.Scan()
			v, _ = strconv.Atoi(scanner.Text())
			v -= 1
			adj[u][v] = 1
		}
	}
	BFS(0)
}

