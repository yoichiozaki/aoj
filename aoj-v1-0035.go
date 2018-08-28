
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func check(x, y, x1, y1, x2, y2 float64) float64 {
	return (x2-x1)*y - (y2-y1)*x + x1*y2 - x2*y1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	a := make([]float64, 8)
	for scanner.Scan() {
		info := strings.Split(scanner.Text(), ",")
		for j := 0; j < 8; j++ {
			a[j], _ = strconv.ParseFloat(info[j], 64)
		}
		if check(a[0], a[1], a[2],a[3], a[6], a[7])*check(a[4], a[5], a[2], a[3], a[6], a[7]) < 0 && check(a[6], a[7], a[0], a[1], a[4], a[5])*check(a[2],a[3], a[0], a[1], a[4], a[5]) < 0 {
			fmt.Println("YES")	
		} else {
			fmt.Println("NO")
		}
	}
}
