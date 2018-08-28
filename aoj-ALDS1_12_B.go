
import (
	"strconv"
	"fmt"
	"bufio"
	"os"
)
const INF = 1 << 21
var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}


func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
	}

	for i := 0;i < n; i++ {
		u := nextInt()
		k := nextInt()
		for j := 0; j < k; j++ {
			v := nextInt()
			c := nextInt()
			adj[u][v] = c
		}
	}

	// for i := range adj {
	// 	fmt.Println(adj[i])
	// }

	dist := make([]int, n)
	prev := make([]int, n)
	visited := make([]bool, n)
	dijkstra := func(start int, size int) {
		for i := 0; i < size; i++ {
			dist[i] = INF
			prev[i] = -1
			visited[i] = false
		}
		dist[start] = 0
		for {
			u := -1
			tmp_min := INF
			for i := 0; i < size; i++ {
				if visited[i] == false && dist[i] < tmp_min {
					tmp_min = dist[i]
					u = i
				}
			}
			if u == -1 {
				break
			}
			visited[u] = true
			for i := 0; i < size; i++ {
				if adj[u][i] > 0 {
					w := adj[u][i]
					newlen := dist[u] + w
					if newlen < dist[i] {
						dist[i] = newlen
						prev[i] = u
					}
				}
			}
		}
	}
	dijkstra(0, n)
	for i, cost := range dist {
		fmt.Printf("%d %d\n", i, cost)
	}
}
