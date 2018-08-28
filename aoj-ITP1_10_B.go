
import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, s, l, h float64
	fmt.Scan(&a, &b, &c)
	s = a * b * math.Sin(math.Pi * c / 180) / 2
	l = a + b + math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) - 2 * a * b *math.Cos(math.Pi * c / 180))
	h = s / a * 2
	fmt.Printf("%f\n", s)
	fmt.Printf("%f\n", l)
	fmt.Printf("%f\n", h)
}
