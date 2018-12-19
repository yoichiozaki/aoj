
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	dict := make(map[string]int)
	dict[" "] = 0
	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10
	dict["L"] = 50
	dict["C"] = 100
	dict["D"] = 500
	dict["M"] = 1000
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		line = append(line, " ")
		sum := 0
		for i := 0; i < len(line)-1; i++ {
			if dict[line[i]] < dict[line[i+1]] {
				sum -= dict[line[i]]
			} else {
				sum += dict[line[i]]
			}
		}
		fmt.Println(sum)
	}
}
