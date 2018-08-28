
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(buf[0])
		b, _ := strconv.Atoi(buf[1])
		n, _ := strconv.Atoi(buf[2])
		sum := 0
		a = 10*(a%b)
		for i := 0; i < n; i++ {
			sum += a/b
			a = 10*(a%b)
		}
		fmt.Printf("%d\n", sum)
	}
}
