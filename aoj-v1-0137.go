
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		seed, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("Case %d:\n", i+1)
		for j := 0; j < 10; j++ {
			tmp := fmt.Sprintf("%08d", seed*seed)
			seed, _ = strconv.Atoi(tmp[2:6])
			fmt.Println(seed)
		}
	}
}
