
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		var sum int
		for _, v := range row {
			num, _ := strconv.Atoi(v)
			sum += num
		}
		fmt.Println(len([]rune(strconv.Itoa(sum))))
	}
}
