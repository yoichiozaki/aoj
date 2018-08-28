
import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

const N = 100

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, N+1)
	m := make([][]int, N+1)
	for i := range m {
		m[i] = make([]int, N+1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		p[i-1], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := 1; i <= n; i++ {
		m[i][i] = 0
	}
	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i+l-1
			m[i][j] = 1 << 21
			for k := i; k <= j-1; k++ {
				m[i][j] = min(m[i][j], m[i][k]+m[k+1][j]+p[i-1]*p[k]*p[j])
			}
		}
	}
	fmt.Println(m[1][n])
}
