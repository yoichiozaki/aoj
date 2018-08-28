
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
		input := strings.Split(scanner.Text(), " ")
		// log.Println(input)
		for i := 0; i < len(input); i++ {
			if strings.Contains(input[i], "Hoshino") {
				input[i] = strings.Replace(input[i], "Hoshino", "Hoshina", -1)
			}
		}
		fmt.Println(strings.Join(input, " "))
	}
}
