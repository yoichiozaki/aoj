
import (
	"bufio"
	"os"
	"unicode"
	"strconv"
	"fmt"
	"strings"
	)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := strings.Split(scanner.Text(), "")
		ans, dir := make([]string, 0), ">"
		for _, x := range s {
			if IsLetter(x) {
				if dir == ">" {
					if !isIn(ans, x) {
						ans = append(ans, x)
					}
				} else {
					if !isIn(ans, x) {
						ans = append([]string{x}, ans...)
					}
				}
			} else if x == ">" || x == "<" {
				dir = x
			}
		}
		fmt.Println(strings.Join(ans, ""))
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isIn(arr []string, target string) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}

