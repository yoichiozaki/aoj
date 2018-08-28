
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
		input := strings.Split(scanner.Text(), ",")
		ele := make([]float64, 0)
		for _, v := range input {
			tmp, _ := strconv.ParseFloat(v, 64)
			ele = append(ele, tmp)
		}
		points := make([]complex128, 0)
		points = append(points, complex(ele[0], ele[1]))
		points = append(points, complex(ele[2], ele[3]))
		target := complex(ele[4], ele[5])
		ans := solver(points, target)
		fmt.Printf("%f %f\n", real(ans), imag(ans))
	}
}
func projection(points []complex128, target complex128) complex128 {
	base := points[1] - points[0]
	r := dot(target-points[0], base) / norm(base)
	base = complex(real(base)*r, imag(base)*r)
	return points[0] + base
}

func solver (points []complex128, target complex128) complex128 {
	return target + complex(real(projection(points, target) - target)*2, imag(projection(points, target) - target)*2)
}

func dot(a, b complex128) float64 {
	return real(a)*real(b) + imag(a)*imag(b)
}

func norm(a complex128) float64 {
	return real(a)*real(a) + imag(a)*imag(a)
}
