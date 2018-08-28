
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
		buf := strings.Split(scanner.Text(), "")
		// log.Println(buf)
		result := make([]string, 0)
		for i := 0; i < len(buf); i++ {
			if buf[i] == "@" {
				repeated := buf[i+2]
				time, _ := strconv.Atoi(buf[i+1])
				for j := 0; j < time; j++ {
					result = append(result, repeated)
				}
				i += 2
				// log.Println(result)
			} else {
				result = append(result, buf[i])
				// log.Println(result)
			}
		}
		fmt.Println(strings.Join(result, ""))
	}
}
