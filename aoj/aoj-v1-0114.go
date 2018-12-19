
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
		input := strings.Split(scanner.Text() , " ")
		a1, _ := strconv.Atoi(input[0])
		m1, _ := strconv.Atoi(input[1])
		a2, _ := strconv.Atoi(input[2])
		m2, _ := strconv.Atoi(input[3])
		a3, _ := strconv.Atoi(input[4])
		m3, _ := strconv.Atoi(input[5])
		if a1 == 0 {
			break
		}
		n := 0
		b1, b2, b3 := 1, 1, 1
		a, b, c := 0, 0, 0
		for {
			if a == 0 {
				b1 = (b1*a1) % m1
			}
			if b == 0 {
				b2 = (b2 * a2) % m2
			}
			if c == 0 {
				b3 = (b3 * a3) % m3
			}
			n++
			if a == 0 && b1 == 1 {
				a = n
			}
			if b == 0 && b2 == 1 {
				b = n
			}
			if c == 0 && b3 == 1 {
				c = n
			}
			if a != 0 && b != 0 && c != 0 {
				break
			}
		}
		ans := a / gcd(a, b) * b
		ans = ans / gcd(ans, c) * c
		fmt.Println(ans)
	}
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		r := a % b
		a, b = b, r
	}
	return a
}

