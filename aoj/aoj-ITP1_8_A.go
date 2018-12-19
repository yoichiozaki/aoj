
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	sr := ""
	for _, x := range s {
		xs := string(x)
		// fmt.Println(xs)
		u := strings.ToUpper(xs)
		l := strings.ToLower(xs)
		if u == xs {
			sr += l
		} else if l == xs {
			sr += u
		} else {
			sr += xs
		}
	}
	fmt.Println(sr)
}
