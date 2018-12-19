
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	text := strings.Split(scanner.Text(), " ")
	ans := make([]string, 0)
	for i, _ := range text {
		if strings.Contains(text[i], ",") {
			text[i] = strings.Replace(text[i], ",", "", -1)
		}
		if strings.Contains(text[i], ".") {
			text[i] = strings.Replace(text[i], ".", "", -1)
		}
		if len(text[i]) >= 3 && len(text[i]) <= 6 {
			ans = append(ans, text[i])
		}
	}
	for i := range ans {
		if i == len(ans) - 1 {
			fmt.Printf("%s\n", ans[i])
		} else {
			fmt.Printf("%s ", ans[i])
		}
	}
}
