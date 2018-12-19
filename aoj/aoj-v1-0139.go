
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
		ans := "NA"
		scanner.Scan()
		s := scanner.Text()
		if s[:2]==">'" && s[len(s)-1:] == "~" {
			tmp := s[2:len(s)-1]
			t := strings.Split(tmp, "#")
			if len(t) == 2 && len(t[0]) > 0 {
				valid := true
				for _, v := range t[0] {
					if string(v) != "=" {
						valid = false
						break
					}
				}
				if valid && t[0] == t[1] {
					ans = "A"
				}
			}
		} else if s[:2] == ">^" && s[len(s)-2:] == "~~" {
			t := s[2:len(s)-2]
			if len(t) > 0 && len(t)%2 == 0 {
				valid := true
				for i := 0; i < len(t); i += 2 {
					if t[i:i+2] != "Q=" {
						valid = false
						break
					}
				}
				if valid {
					ans = "B"
				}
			}
 		}
		fmt.Println(ans)
	}
}
