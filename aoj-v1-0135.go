
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), ":")
		h, _ := strconv.Atoi(input[0])
		m, _ := strconv.Atoi(input[1])
		H, M := (30*h+(m/2))*2, (6*m)*2
		if m&1 == 1 {
			H++
		}
		a := H - M
		if a < 0 {
			a = -a
		}
		a2 := 720 - a
		if a2 < a {
			a = a2
		}
		if a < 60 {
			fmt.Println("alert")
		} else if a >= 180 && a <= 360 {
			fmt.Println("safe")
		} else {
			fmt.Println("warning")
		}
	}
}
