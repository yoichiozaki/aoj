
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
	"fmt"
	)

const esp = 1.0e-8

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		circles := make([]complex128, 0)
		for i := 0; i< n; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), ",")
			// log.Println(input)
			x, _ := strconv.ParseFloat(input[0], 64)
			y, _ := strconv.ParseFloat(input[1], 64)
			circles = append(circles, complex(x, y))
		}
		// log.Println(circles)
		cross := make([]complex128, 0)
		for i := 0; i < n; i++ {
			for j := i+1; j < n; j++ {
				crossPoint(circles[i], circles[j], &cross)
			}
		}
		// log.Println(cross)
		ans := 0
		for _, i := range cross {
			f := 0
			for _, j := range circles {
				dx := real(j) - real(i)
				dy := imag(j) - imag(i)
				d := dx*dx + dy*dy
				if math.Abs(d - 1.0) <= esp || d <= 1.0 {
					f++
				}
			}
			if f > ans {
				ans = f
			}
		}
		if ans == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(ans)
		}
	}
}

func crossPoint(p1, p2 complex128, cross *[]complex128) {
	a := (real(p1) - real(p2)) / 2
	b := (imag(p1) - imag(p2)) / 2
	d := math.Sqrt(a*a + b*b)
	if d*d > 1.0 + esp {
		// log.Printf("d*d = %f", d*d)
		return
	}
	h := math.Sqrt(1 - d*d)
	k, X, Y := h/d, (real(p1) + real(p2)) / 2, (imag(p1) + imag(p2)) / 2
	*cross = append(*cross, complex(X+k*b, Y-k*a))
	*cross = append(*cross, complex(X-k*b, Y+k*a))
	// log.Println(*cross)
}
