
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	to := make([][]int, m)
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(input[0])
		y, _ := strconv.Atoi(input[1])
		to[x-1] = append(to[x-1], y-1)
	}
	topologicalSort(m, to)
}

func topologicalSort(V int, to [][]int) {
	cnt := make([]int, V)
	for i := 0; i < V; i++ {
		for _, j := range to[i] {
			cnt[j]++
		}
	}
	Q := make([]int, 0)
	for i := 0; i < V; i++ {
		if cnt[i] == 0 {
			Q = append(Q, i)
		}
	}
	for len(Q) > 0 {
		i := Q[0]
		Q = Q[1:]
		fmt.Println(i+1)
		for _, k := range to[i] {
			cnt[k]--
			if cnt[k] == 0 {
				Q = append(Q, k)
			}
		}
	}
}
