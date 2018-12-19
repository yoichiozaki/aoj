
import (
	"fmt"
	"math"
)

func main() {
	n := scanInt()
	for i := 0; i < n; i++ {
		cal()
	}
}

func scanInt() int {
	var n int
	fmt.Scanf("%d", &n)
	return n
}

func cal() {
	var x1, y1, x2, y2, x3, y3 float64
	fmt.Scanf("%f%f%f%f%f%f", &x1, &y1, &x2, &y2, &x3, &y3)
	a := 2 * (x1 - x2)
	b := 2 * (y1 - y2)
	c := 2 * (x1 - x3)
	d := 2 * (y1 - y3)
	e := x1*x1 - x2*x2 + y1*y1 - y2*y2
	f := x1*x1 - x3*x3 + y1*y1 - y3*y3

	px := (d*e - b*f) / (a*d - b*c)
	py := (a*f - c*e) / (a*d - b*c)

	r := math.Sqrt((x1-px)*(x1-px) + (y1-py)*(y1-py))
	fmt.Printf("%.3f %.3f %.3f\n", px, py, r)
}

