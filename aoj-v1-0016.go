
import (
	"fmt"
	"math"
	"strings"
	"strconv"
)

func main() {
	var a, b float64
	var x, y float64
	x = 0
	y = 0
	var now float64
	now = 90
	var info string
	for {
		fmt.Scan(&info)
		infos := strings.Split(info, ",")
		a_int, _ := strconv.Atoi(infos[0])
		b_int, _ := strconv.Atoi(infos[1])
		a = float64(a_int)
		b = float64(b_int)
		// println(a, b)
		if a == 0 && b == 0 {
			break
		}
		x += math.Cos(now*math.Pi/180)*a
		y += math.Sin(now*math.Pi/180)*a
		now -= b
	}
	fmt.Printf("%d\n%d\n", int(x), int(y))
}
