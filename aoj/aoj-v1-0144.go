
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
const INF = 1<<31

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		buf := strings.Split(scanner.Text(), " ")
		input := make([]int, len(buf))
		for j := range input {
			input[j], _ = strconv.Atoi(buf[j])
		}
		x, k := input[0]-1, input[1]
		for j := 0; j < k; j++ {
			y := input[j+2]-1
			arr[x][y] = 1
		}
	}
	for k := 0; k < n; k++ {
		arr[k][k] = 0
		for i := 0; i < n; i++ {
			if arr[i][k] >= INF {
				continue
			}
			for j := 0; j < n; j++ {
				if arr[k][j] >= INF {
					continue
				}
				arr[i][j] = min(arr[i][j], arr[i][k]+arr[k][j])
			}
		}
	}
	scanner.Scan()
	p, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < p; i++ {
		scanner.Scan()
		buf := strings.Split(scanner.Text(), " ")
		input := make([]int, len(buf))
		for j := range input {
			input[j], _ = strconv.Atoi(buf[j])
		}
		s, d, v := input[0]-1, input[1]-1, input[2]
		if arr[s][d] < v {
			fmt.Println(arr[s][d]+1)
		} else {
			fmt.Println("NA")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
