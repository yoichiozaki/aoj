
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		a, _ := strconv.ParseFloat(scanner.Text(), 64)
		b := a
		c := a
		for i := 2; i <= 10; i++ {
			if i % 2 != 0 {
				b /= 3
			} else {
				b *= 2
			}
			c += b
		}
		fmt.Println(c)
	}
}
