
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
const MAX = 1000

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n > 2*MAX {
			n = 4*MAX - n
		}
		k := n+1
		ans := k*(k+1)*(k+2)/6
		if n > MAX {
			k = n - MAX
			ans -= 2*k*(k+1)*(k+2)/3
		}
		fmt.Println(ans)
	}
}
