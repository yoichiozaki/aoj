
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
		input := strings.Split(scanner.Text() , " ")
		p, _ := strconv.Atoi(input[0])
		q, _ := strconv.Atoi(input[1])
		rr := make([]int, q)
		for i := range rr {
			rr[i] = -1
		}
		rr[p] = 0
		tmpans := make([]string, 0)
		for k := 1; k < q; k++ {
			p *= 10
			tmpans = append(tmpans, strconv.Itoa(p/q))
			r := p%q
			if r == 0 || rr[r] >= 0 {
				fmt.Println(strings.Join(tmpans, ""))
				if rr[r] >= 0 {
					fmt.Println(strings.Repeat(" ", rr[r]) + strings.Repeat("^", k-rr[r]))
				}
				break
			}
			rr[r], p = k, r
		}
	}
}

