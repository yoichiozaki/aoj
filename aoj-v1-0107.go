
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
		d, _ := strconv.Atoi(input[0])
		w, _ := strconv.Atoi(input[1])
		h, _ := strconv.Atoi(input[2])
		if d == 0 && w == 0 && h == 0 {
			break
		}
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		vmin := min(d*d+w*w, d*d+h*h, w*w+h*h)
		for i := 0; i < n; i++ {
			scanner.Scan()
			r, _ := strconv.Atoi(scanner.Text())
			if (2*r)*(2*r) > vmin {
				fmt.Println("OK")
			} else {
				fmt.Println("NA")
			}
		}
	}
}

func min(v1 int, vn ...int) (m int) {
	m = v1
	for i := 0; i < len(vn); i++ {
		if vn[i] < m {
			m = vn[i]
		}
	}
	return
}
