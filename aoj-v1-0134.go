
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
	sum := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmp, _ := strconv.Atoi(scanner.Text())
		sum += tmp
	}
	fmt.Println(sum/n)
}
