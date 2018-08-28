
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func solve(a, b, c int) int {
	if a == 1 && b >= c && b < 10 {
		return 1
	}
	if a == 1 {
		return 0
	}
	ret := 0
	for i := c; i < 10 && i < b; i++ {
		ret += solve(a-1, b-i, i+1)
	}
	return ret
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		fmt.Println(solve(a, b, 0))
	}
}
