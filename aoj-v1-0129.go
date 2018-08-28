
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	x := make([]int, 100)
	y := make([]int, 100)
	r := make([]int, 100)
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		for i := 0; i < a; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			tmpx, _ := strconv.Atoi(input[0])
			tmpy, _ := strconv.Atoi(input[1])
			tmpr, _ := strconv.Atoi(input[2])
			x[i], y[i], r[i] = tmpx, tmpy, tmpr
		}
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < b; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			c, _ := strconv.Atoi(input[0])
			d, _ := strconv.Atoi(input[1])
			e, _ := strconv.Atoi(input[2])
			f, _ := strconv.Atoi(input[3])
			ok := false
			for j := 0; j < a; j++ {
				if (c-x[j])*(c-x[j])+(d-y[j])*(d-y[j]) >= r[j]*r[j] || (e-x[j])*(e-x[j])+(f-y[j])*(f-y[j]) >= r[j]*r[j] {
					if (e-c)*(x[j]-c)+(f-d)*(y[j]-d) < 0 {
						if (x[j]-c)*(x[j]-c)+(y[j]-d)*(y[j]-d) <= r[j]*r[j] {
							ok = true
						}
					} else {
						if (e-c)*(x[j]-c)+(f-d)*(y[j]-d)>(e-c)*(e-c)+(f-d)*(f-d) {
							if (x[j]-e)*(x[j]-e)+(y[j]-f)*(y[j]-f)<=r[j]*r[j] {
								ok = true
							}
						} else {
							if (x[j]-c)*(x[j]-c)+(y[j]-d)*(y[j]-d)-((e-c)*(x[j]-c)+(f-d)*(y[j]-d))*((e-c)*(x[j]-c)+(f-d)*(y[j]-d))/((e-c)*(e-c)+(f-d)*(f-d))<=r[j]*r[j] {
								ok = true
							}
						}
					}

				}
			}
			if ok {
				fmt.Println("Safe")
			} else {
				fmt.Println("Danger")
			}
		}
	}
}

