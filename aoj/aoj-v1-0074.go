
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
		input := strings.Split(scanner.Text(), " ")
		t, _ := strconv.Atoi(input[0])
		h, _ := strconv.Atoi(input[1])
		s, _ := strconv.Atoi(input[2])
		if t == -1 && h == -1 && s == -1 {
			break
		}
		recorded := t*3600 + h*60 + s
		left := 7200 - recorded
		leftx3 := left*3
		fmt.Printf("%02d:%02d:%02d\n%02d:%02d:%02d\n",
			left/3600, left%3600/60, left%60,
			leftx3/3600, leftx3%3600/60, leftx3%60)
	}
}
