
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner= bufio.NewScanner(os.Stdin)

func main() {
	result := make([][]int, 4)
	for i :=range result {
		result[i] = make([]int, 2)
	}
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		l, _ := strconv.ParseFloat(buf[0], 64)
		r, _ := strconv.ParseFloat(buf[1], 64)
		if 1.1 <= l {
			result[0][0]++
		} else if 0.6 <= l && l < 1.1 {
			result[1][0]++
		} else if 0.2 <= l && l < 0.6 {
			result[2][0]++
		} else {
			result[3][0]++
		}
		if 1.1 <= r {
			result[0][1]++
		} else if 0.6 <= r && r < 1.1 {
			result[1][1]++
		} else if 0.2 <= r && r < 0.6 {
			result[2][1]++
		} else {
			result[3][1]++
		}
	}
	ans := make([][]string, 4)
	for i :=range ans {
		ans[i] = make([]string, 2)
	}
	for i := range result {
		for j := range result[i] {
			ans[i][j] = strconv.Itoa(result[i][j])
		}
	}
	for i := 0; i < 4; i++ {
		fmt.Println(strings.Join(ans[i], " "))
	}
}
