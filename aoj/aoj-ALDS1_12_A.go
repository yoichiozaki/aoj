
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const (
	MAX = 100
	INF = 1 << 21
	WHITE = 0
	GRAY = 1
	BLACK = 2
)

var (
	n int
	M [MAX][MAX]int
)

func prim() int {
	var u, minv int
	var d [MAX]int
	var p [MAX]int
	var color [MAX]int

	for i := 0; i < n; i++ {
		d[i] = INF
		p[i] = -1
		color[i] = WHITE
	}
	d[0] = 0
	for {
		minv = INF
		u = -1
		for i := 0; i < n; i++ {
			if minv > d[i] && color[i] != BLACK {
				u = i
				minv = d[i]
			}
		}
		if u == -1 {
			break
		}
		color[u] = BLACK
		for v := 0; v < n; v++ {
			if color[v] != BLACK && M[u][v] != INF {
				if d[v] > M[u][v] {
					d[v] = M[u][v]
					p[v] = u
					color[v] = GRAY
				}
			}
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		if p[i] != -1 {
			sum += M[i][p[i]]
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			scanner.Scan()
			e, _ := strconv.Atoi(scanner.Text())
			if e == -1 {
				M[i][j] = INF
			} else {
				M[i][j] = e
			}
		}
	}
	fmt.Println(prim())
}
