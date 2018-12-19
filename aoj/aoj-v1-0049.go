
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	A, B, O, AB := 0, 0, 0, 0
	for scanner.Scan() {
		meminfo := strings.Split(scanner.Text(), ",")
		if meminfo[1] == "A" {
			A++
		}
		if meminfo[1] == "B" {
			B++
		}
		if meminfo[1] == "O" {
			O++
		}
		if meminfo[1] == "AB" {
			AB++
		}
	}
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(AB)
	fmt.Println(O)
}
