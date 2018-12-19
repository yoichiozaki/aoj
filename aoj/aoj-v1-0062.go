
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
		inputs := strings.Split(scanner.Text(), "")
		var numbers [10]int
		for i := 0; i < 10; i++ {
			numbers[i], _ = strconv.Atoi(inputs[i])
		}
		// fmt.Println(numbers)
		for i := 9; i > 0; i-- {
			for j := 0; j < i; j++ {
				numbers[j] = (numbers[j] + numbers[j+1]) % 10
			}
		}
		// fmt.Println(numbers)
		fmt.Println(numbers[0])
	}
}
