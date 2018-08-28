
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
const tubo = 3.305785

func main() {
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	a, _ := strconv.ParseFloat(input[0], 64)
	b, _ := strconv.ParseFloat(input[1], 64)
	fmt.Println(a*b/tubo)
}
