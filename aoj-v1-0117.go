
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const MAX = 99999999

func main() {
	g := make([][]int, 20)
	for i := range g {
		g[i] = make([]int, 20)
	}
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = MAX
		}
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), ",")
		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])
		c, _ := strconv.Atoi(input[2])
		d, _ := strconv.Atoi(input[3])
		g[a-1][b-1] = c
		g[b-1][a-1] = d
	}
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")
	s, _ := strconv.Atoi(input[0])
	gg, _ := strconv.Atoi(input[1])
	V, _ := strconv.Atoi(input[2])
	P, _ := strconv.Atoi(input[3])
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if g[i][j] > g[i][k]+g[k][j] {
					g[i][j] = g[i][k] + g[k][j]
				}
			}
		}
	}
	fmt.Println(V - P - g[s-1][gg-1] - g[gg-1][s-1])
}

