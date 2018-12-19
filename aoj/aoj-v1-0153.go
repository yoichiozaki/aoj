
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		buf0 := strings.Split(scanner.Text(), " ")
		input0 := make([]int, 2)
		for i := range input0 {
			input0[i], _ = strconv.Atoi(buf0[i])
		}
		if input0[0] == 0 && input0[1] == 0 {
			break
		}
		a, b := input0[0], input0[1]
		scanner.Scan()
		buf1 := strings.Split(scanner.Text(), " ")
		input1 := make([]int, 2)
		for i := range input1 {
			input1[i], _ = strconv.Atoi(buf1[i])
		}
		c, d := input1[0], input1[1]
		scanner.Scan()
		buf2 := strings.Split(scanner.Text(), " ")
		input2 := make([]int, 2)
		for i := range input2 {
			input2[i], _ = strconv.Atoi(buf2[i])
		}
		e, f := input2[0], input2[1]
		scanner.Scan()
		buf3 := strings.Split(scanner.Text(), " ")
		input3 := make([]int, 2)
		for i := range input3 {
			input3[i], _ = strconv.Atoi(buf3[i])
		}
		g, h := input3[0], input3[1]
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())
		if (a-g)*(a-g)+(b-h)*(b-h) <= r*r && (c-g)*(c-g)+(d-h)*(d-h) <= r*r && (e-g)*(e-g)+(f-h)*(f-h) <= r*r {
			fmt.Println("b")
		} else if dist(a, b, c, d, g, h) > float64(r)-0.000001 && dist(a, b, e, f, g, h) > float64(r)-0.000001 && dist(e, f, c, d, g, h) > float64(r)-0.000001 && inclusive(a, b, c, d, e, f, g, h) {
			fmt.Println("a")
		} else if CL_intersect(a, b, c, d, g, h, r) || CL_intersect(a, b, e, f, g, h, r) || CL_intersect(e, f, c, d, g, h, r) {
			fmt.Println("c")
		} else {
			fmt.Println("d")
		}
	}
}

func CL_intersect(ax, ay, bx, by, cx, cy, r int) bool {
	acx := cx - ax
	acy := cy - ay
	bcx := cx - bx
	bcy := cy - by
	abx := bx - ax
	aby := by - ay
	if acx*acx+acy*acy >= r*r || bcx*bcx+bcy*bcy >= r*r {
		if abx*acx+aby*acy < 0 {
			if acx*acx+acy*acy <= r*r {
				return true
			}
		} else {
			if abx*acx+aby*acy > abx*abx+aby*aby {
				if bcx*bcx+bcy*bcy <= r*r {
					return true
				}
			} else if acx*acx+acy*acy-(abx*acx+aby*acy)*(abx*acx+aby*acy)/(abx*abx+aby*aby) <= r*r {
				return true
			}
		}
	}
	return false
}

func inclusive(ax, ay, bx, by, cx, cy, px, py int) bool {
	ok1 := true
	ok2 := true
	if (bx-ax)*(py-ay)-(px-ax)*(by-ay) < 0 {
		ok1 = false
	}
	if (bx-ax)*(py-ay)-(px-ax)*(by-ay) > 0 {
		ok2 = false
	}
	if (cx-bx)*(py-by)-(px-bx)*(cy-by) < 0 {
		ok1 = false
	}
	if (cx-bx)*(py-by)-(px-bx)*(cy-by) > 0 {
		ok2 = false
	}
	if (ax-cx)*(py-cy)-(px-cx)*(ay-cy) < 0 {
		ok1 = false
	}
	if (ax-cx)*(py-cy)-(px-cx)*(ay-cy) > 0 {
		ok2 = false
	}
	return ok1 || ok2
}

func dist(ax, ay, bx, by, px, py int) float64 {
	dx := float64(bx - ax)
	dy := float64(by - ay)
	A := float64(dx*dx + dy*dy)
	B := dx*float64(ax-px) + dy*float64(ay-py)
	t := -B / A
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	tx := float64(ax) + dx*t
	ty := float64(ay) + dy*t
	return math.Sqrt((float64(px)-tx)*(float64(px)-tx) + (float64(py)-ty)*(float64(py)-ty))
}

