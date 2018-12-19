
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		n /= 100
		ret := 99999999
		for i := 0; i <= 25; i++ {
			for j := 0; j <= 16; j++ {
				for k := 0; k <= 10; k++ {
					if i*2+j*3+k*5 == n {
						col := 0
						col += ((i / 5) * 380 * 4)
						col += ((i % 5) * 380)
						col += ((j / 4) * 550 * 4 * 17 / 20)
						col += ((j % 4) * 550)
						col += ((k / 3) * 850 * 3 * 22 / 25)
						col += ((k % 3) * 850)
						ret = min(ret, col)
					}
				}
			}
		}
		fmt.Println(ret)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

