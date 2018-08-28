
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	dp := make([][]int, 11)
	for i := range dp {
		dp[i] = make([]int, 1001)
	}
	dp[0][0] = 1
	for k := 0; k < 101; k++ {
		for n := 9; n > -1; n-- {
			for s := 0; s < 1001-k; s++ {
				dp[n+1][s+k] += dp[n][s]
			}
		}
	}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(input[0])
		s, _ :=strconv.Atoi(input[1])
		if n == 0 && s == 0 {
			break
		}
		fmt.Println(dp[n][s])
	}
}
