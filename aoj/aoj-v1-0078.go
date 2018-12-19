
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		table := make([][]int, n)
		for i := 0; i < n; i++ {
			table[i] = make([]int, n)
			for j := 0; j < n; j++ {
				table[i][j] = 0
			}
		}
		if n == 0 {
			break
		}
		row := n/2 + 1
		col := n/2
		for i := 0; i < n*n; i++ {
			if table[row%n][col%n] != 0 {
				row++
				col--
			}
			table[row%n][col%n] = i+1
			row++
			col++
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%4d", table[i][j])
			}
			fmt.Println()
		}
	}
}
