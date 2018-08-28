
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
	"sort"
)

var scanner= bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		mp := make([]string, n)
		for i := range mp {
			scanner.Scan()
			mp[i] = "0" + scanner.Text() + "0"
		}
		tmp := make([]string, n+2)
		for i := range tmp {
			tmp[i] = "0"
		}
		mp = append([]string{strings.Join(tmp, "")}, mp...)
		mp = append(mp, strings.Join(tmp, ""))
		score := make([][][]int, n+2)
		for i := range score {
			score[i] = make([][]int, n+2)
			for j := range score[i] {
				score[i][j] = make([]int, 4)
			}
		}
		max_score := 0
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				tmps := mp[i]
				if string(tmps[j]) == "1" {
					score[i][j][0] = score[i - 1][j][0] + 1
					score[i][j][1] = score[i][j - 1][1] + 1
					score[i][j][2] = score[i - 1][j - 1][2] + 1
					score[i][j][3] = score[i - 1][j + 1][3] + 1
					max_score = max(max_score, score[i][j][0], score[i][j][1], score[i][j][2], score[i][j][3])
				}
			}
		}
		fmt.Println(max_score)
	}
}

func max(a, b, c, d, e int) int {
	buf := []int{a, b, c, d, e}
	sort.Ints(buf)
	return buf[len(buf)-1]
}
