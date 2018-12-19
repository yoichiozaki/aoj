
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

const MAXINT = 1<<32

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		check_table := make([]int, 100)
		table := make([][]int, 100)
		fix := make([]int, 100)
		for i := range table {
			check_table[i] = 0
			fix[i] = 0
			table[i] = make([]int, 100)
			for j := range table[i] {
				table[i][j] = 0
			}
		}
		if scanner.Text() == "0" {
			break
		}
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < m; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), ",")
			a, _ := strconv.Atoi(input[0])
			b, _ := strconv.Atoi(input[1])
			d, _ := strconv.Atoi(input[2])
			table[a][b] = d
			table[b][a] = d
		}
		check_table[0] = 1
		fix[0] = 0
		count := 1
		num := 0
		min_index := 0
		for count < n {
			tmp_min := MAXINT
			for i := 0; i < count; i++ {
				for k := 0; k < n; k++ {
					if table[fix[i]][k] > 0 && check_table[k] == 0 && table[fix[i]][k] < tmp_min {
						tmp_min = table[fix[i]][k]
						min_index = k
					}
				}
			}
			num += tmp_min/100 - 1
			check_table[min_index] = 1
			fix[count] = min_index
			count++
		}
		fmt.Println(num)
	}
}
