
import (
	"bufio"
			"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		if m == 0 {
			break
		}
		ab := make([][]int, 0)
		for i := 0; i < m; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			buf := make([]int, len(input))
			for j := range buf {
				buf[j], _ = strconv.Atoi(input[j])
			}
			ab = append(ab, buf)
		}
		dp := make([][]int, 8)
		for i := range dp {
			dp[i] = make([]int, 1001)
		}
		dp[0][0] = 1
		for i := 1; i < m+1; i++ {
			for j := 0; j < 1001; j++ {
				if dp[i-1][j] == 0 {
					continue
				}
				s := j
				for k := ab[i-1][1]; k > -1; k-- {
					if s < 1001 {
						dp[i][s] += dp[i-1][j]
					}
					s += ab[i-1][0]
				}
			}
		}
		scanner.Scan()
		g, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < g; i++ {
			scanner.Scan()
			n, _ := strconv.Atoi(scanner.Text())
			fmt.Println(dp[m][n])
		}
	}
}

