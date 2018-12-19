
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
		H, _ := strconv.Atoi(input[0])
		W, _ := strconv.Atoi(input[1])
		if H == 0 && W == 0 {
			break
		}
		a := make([][]string, H)
		for r := 0; r < H; r++ {
			scanner.Scan()
			a[r] = strings.Split(scanner.Text(), "")
		}
		r, c := 0, 0
		for {
			k := a[r][c]
			a[r][c] = "#"
			if k == "#" {
				// Loop when you reach the place you have already visited
				fmt.Println("LOOP")
				break
			} else if k == ">" {
				c++
			} else if k == "<" {
				c--
			} else if k == "^" {
				r--
			} else if k == "v" {
				r++
			} else if k == "." {
				fmt.Printf("%d %d\n", c, r)
				break
			}
		}
	}
}
