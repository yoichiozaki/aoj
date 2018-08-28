
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		s := scanner.Text()
		if s == "0" {
			break
		}
		sum := 0
		for _, c := range s {
			sum += int(c - '0')
		}
		fmt.Println(sum)
	}
}
