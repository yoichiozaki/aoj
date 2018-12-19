
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		first := true
		beki := 1
		for num != 0 {
			if num % 2 == 1 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(beki)
				first = false
			}
			num /= 2
			beki *= 2
		}
		fmt.Println()
	}
}
