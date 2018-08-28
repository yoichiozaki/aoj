
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
		input := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(input[0])
		m, _ := strconv.Atoi(input[1])
		if n == 0 && m == 0 {
			break
		}
		k := 1
		for j := 2; j < n+1; j++ {
			k = (k+m)%j
			if k == 0 {
				k = j
			}
		}
		fmt.Println(k)
	}
}
