
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	a := true
	b, c := false, false
	for scanner.Scan() {
		var d bool
		do := strings.Split(scanner.Text(), ",")
		x := do[0]
		y := do[1]
		if x == "A" && y == "B" || x == "B" && y == "A" {
			d = a
			a = b
			b = d
		} else if x == "B" && y == "C" || x == "C" && y == "B" {
			d = b
			b = c
			c = d
		} else {
			d = c
			c = a
			a = d
		}
	}
	if a {
		fmt.Printf("A\n")
	}
	if b {
		fmt.Printf("B\n")
	}
	if c {
		fmt.Printf("C\n")
	}
}
