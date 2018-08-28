
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
func main() {
	scanner.Scan()
	input := scanner.Text()
	input = strings.Replace(input, "apple", "00000", -1)
	input = strings.Replace(input, "peach", "11111", -1)
	input = strings.Replace(input, "00000", "peach", -1)
	input = strings.Replace(input, "11111", "apple", -1)
	fmt.Println(input)
}
