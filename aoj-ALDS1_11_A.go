
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	Adj := make([][]int, n)
	for i := range Adj {
		Adj[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		if k == 0 {
			continue
		}
		for i := 0; i < k; i++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			Adj[u-1][v-1] = 1
		}
	}
	for i := range Adj {
		for j := range Adj[i] {
			if j == n-1 {
				fmt.Printf("%d\n", Adj[i][j])
				continue
			}
			fmt.Printf("%d ", Adj[i][j])
		}
	}
}
