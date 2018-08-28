
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
const INF = 0x7fffffff
func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	top := make([]int, 0)
	bottom := make([]int, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		buf := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(buf[0])
		b, _ := strconv.Atoi(buf[1])
		top = append(top, a)
		bottom = append(bottom, b)
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = INF
		}
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 0
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			tmp := make([]int, 0)
			for k := 0; k < i; k++ {
				tmp = append(tmp, dp[k][j] + dp[i - k - 1][j + k + 1] +top[j] * bottom[j + k] * top[j + k + 1] * bottom[j + i])
			}
			sort.Ints(tmp)
			dp[i][j] = tmp[0]
		}
	}
	fmt.Println(dp[n-1][0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
