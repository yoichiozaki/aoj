
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	table := make([]int, n+1)
	vmax := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(input[0])
		v, _ := strconv.Atoi(input[1])
		table[a] = v
		if table[a] > vmax {
			vmax = table[a]
		}
	}
	for i := 1; i < n+2; i++ {
		if table[i] == vmax {
			fmt.Printf("%d %d\n", i, vmax)
			break
		}
	}
}
