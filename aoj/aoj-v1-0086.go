
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
var deg [103]int

func main() {
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])
		deg[a]++
		deg[b]++
		for scanner.Scan() {
			input := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(input[0])
			b, _ := strconv.Atoi(input[1])
			if a == 0 {
				break
			}
			deg[a]++
			deg[b]++
		}
		p := 0
		for i := 3; i < 103; i++ {
			p += deg[i]%2
		}
		if p == 0 && deg[2]%2 != 0 && deg[1]%2 != 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
		for i := range deg {
			deg[i] = 0
		}
	}
}
