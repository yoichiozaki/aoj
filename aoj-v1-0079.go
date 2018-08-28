
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	now := 1
	x := make([]float64, 22)
	y := make([]float64, 22)
	for scanner.Scan() {
		// if scanner.Text() == "end" {
		// 	break
		// }
		input := strings.Split(scanner.Text(), ",")
		x[now], _ = strconv.ParseFloat(input[0], 64)
		y[now], _ = strconv.ParseFloat(input[1], 64)
		now++
	}
	// fmt.Println(x)
	// fmt.Println(y)
	x[0] = x[now-1]
	x[now] = x[1]
	var S float64 = 0
	for i := 1; i < now; i++ {
		S += (x[i+1]-x[i-1])*y[i]
	}
	if S < 0 {
		S = -S
	}
	fmt.Println(S/2)
}
