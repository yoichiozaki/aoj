
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
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		buf := make([][]string, n)
		for i := range buf {
			buf[i] = make([]string, n)
		}
		table := make([][]int, 1001)
		for i := range table {
			table[i] = make([]int, 1001)
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf[i] = strings.Split(scanner.Text(), "")
		}
		// log.Println(table)
		max_r := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if buf[i][j] == "." {
					table[i+1][j+1] = min(table[i][j],min(table[i][j+1],table[i+1][j]))+1
					max_r = max(max_r,table[i+1][j+1])
				} else {
					table[i+1][j+1] = 0
				}
			}
		}
		fmt.Println(max_r)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
