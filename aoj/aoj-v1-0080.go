
import (
	"bufio"
	"os"
	"strconv"
	"math"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		q, _ := strconv.ParseFloat(scanner.Text(), 64)
		if q == -1 {
			break
		}
		p := q/2
		for {
			p = p - (p*p*p-q)/(3*p*p)
			if math.Abs(p*p*p-q) < q*0.00001 {
				break
			}
		}
		fmt.Println(p)
	}
}
