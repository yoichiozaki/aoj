
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
		table := make([][]int, n+1)
		for i := range table {
			table[i] = make([]int, n+1)
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			for j := 0; j < n; j++ {
				table[i][j], _ = strconv.Atoi(input[j])
			}
		}
		// log.Println(table)
		for i := 0; i < n; i++ {
			sum := 0
			for j := 0; j < n; j++ {
				sum += table[i][j]
			}
			table[i][n] = sum
		}
		// log.Println(table)
		for i := 0; i < n; i++ {
			sum := 0
			for j := 0; j < n; j++ {
				sum += table[j][i]
			}
			table[n][i] = sum
		}
		// log.Println(table)
		sum := 0
		for i := 0; i < n; i++ {
			sum += table[n][i]
		}
		table[n][n] = sum
		// log.Println(table)
		for i := range table {
			for j, vj := range table[i] {
				if j == n {
					fmt.Printf("%5d\n", vj)
				} else {
					fmt.Printf("%5d", vj)
				}
			}
		}
	}
}
