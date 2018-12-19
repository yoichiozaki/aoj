
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	ESP = 0.0000001
	LOWER = -ESP
	UPPER = 1 + ESP
)

func main() {
	uaz := scanPoint()
	enemy := scanPoint()
	A := scanPoint()
	B := scanPoint()
	C := scanPoint()
	a := sub(enemy, uaz)
	b := sub(A, B)
	c := sub(A, C)
	d := sub(A, uaz)
	denom := det(a, b, c)
	if denom != 0 {
		t := det(d, b, c) / denom
		u := det(a, d, c) / denom
		v := det(a, b, d) / denom
		if t < LOWER {
			fmt.Println("HIT")
		} else if LOWER < t && t < UPPER &&
			LOWER <= u && u <= UPPER &&
			LOWER <= v && v <= UPPER &&
			LOWER <= u + v && u + v <= UPPER {
			fmt.Println("MISS")
		} else {
			fmt.Println("HIT")
		}
	} else {
		fmt.Println("HIT")
	}
}

func scanPoint() [3]float64 {
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	point := [3]float64{}
	for i := 0; i < 3; i++ {
		point[i], _ = strconv.ParseFloat(input[i], 64)
	}
	return point
}

func det(a, b, c [3]float64) float64 {
	return a[0] * b[1] * c[2] +
		a[2] * b[0] * c[1] +
		a[1] * b[2] * c[0] -
		a[2] * b[1] * c[0] -
		a[1] * b[0] * c[2] -
		a[0] * b[2] * c[1]
}

func sub(v0, v1 [3]float64) [3]float64 {
	return [3]float64{v0[0] - v1[0], v0[1] - v1[1], v0[2] - v1[2]}
}

