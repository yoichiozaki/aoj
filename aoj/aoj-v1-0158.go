
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
		if n == 0 {
			break
		}
		k := 0
		for n > 1 {
			if n&1 != 0 {
				n = n+((n+1)>>1)
				k+=2
			} else {
				n >>= 1
				k++
			}
		}
		fmt.Println(k)
	}
}
