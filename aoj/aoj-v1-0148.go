
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner= bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n%39 == 0 {
			fmt.Println(fmt.Sprintf("3C%02d", 39))
			continue
		}
		fmt.Println(fmt.Sprintf("3C%02d", n%39))
	}
}
