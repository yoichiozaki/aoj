
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
		n, _ := strconv.Atoi(scanner.Text())
		if n == -1 {
			break
		}
		var x, y float64
		x = 1
		y = 0
		r := math.Pi/2
		for i := 0; i < n-1; i++ {
			y += math.Sin(r)
			x += math.Cos(r)
			r = math.Atan(y/x)
			if x < 0 {
				r += math.Pi
			}
			r += math.Pi/2
		}
		fmt.Printf("%.2f\n", x)
		fmt.Printf("%.2f\n", y)
	}
}
