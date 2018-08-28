
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		x, _ := strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		h, _ := strconv.ParseFloat(scanner.Text(), 64)
		if x == 0 && h == 0 {
			break
		}
		fmt.Printf("%f\n",2*math.Sqrt(x/2*x/2+h*h)*x+x*x)
	}
}
