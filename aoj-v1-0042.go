
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	Case := 1
	for {
		scanner.Scan()
		W, _ := strconv.Atoi(scanner.Text())
		if W == 0 {
			break
		}
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		w := make([]int, n+1)
		v := make([]int, n+1)
		for i := 1; i <= n; i++ {
			scanner.Scan()
			info := strings.Split(scanner.Text(), ",")
			v[i], _ = strconv.Atoi(info[0])
			w[i], _ = strconv.Atoi(info[1])
		}
		// fmt.Println(v)
		// fmt.Println(w)
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, W+1)
		}
		maxV := -(1 << 21)
		minW := 1 << 21
		for i := 1; i <= n; i++ {
			for j := 1; j <= W; j++ {
				// fmt.Println(dp)
				if j-w[i] >= 0 {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
					if maxV < dp[i][j] {
						maxV = dp[i][j]
						minW = j
					} else if maxV == dp[i][j] {
						minW = min(minW, j)
					}
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
		fmt.Printf("Case %d:\n", Case)
		Case += 1
		fmt.Println(maxV)
		fmt.Println(minW)
	}
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
