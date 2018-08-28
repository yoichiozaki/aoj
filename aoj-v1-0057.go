
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		ans := 1
		n, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < n; i++ {
			ans += i+1
		}
		fmt.Println(ans)
	}
}
