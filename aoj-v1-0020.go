
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToUpper(input)
	fmt.Println(input)
}
