
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		buf := make([]float64, len(input))
		for j := range buf {
			buf[j], _ = strconv.ParseFloat(input[j], 64)
		}
		p1 := complex(buf[0], buf[1])
		p2 := complex(buf[2], buf[3])
		p3 := complex(buf[4], buf[5])
		pk := complex(buf[6], buf[7])
		ps := complex(buf[8], buf[9])
		kin := isAtSameSide(p3, pk, []complex128{p1, p2}) &&
			isAtSameSide(p1, pk, []complex128{p2, p3}) &&
			isAtSameSide(p2, pk, []complex128{p3, p1})
		sin := isAtSameSide(p3, ps, []complex128{p1, p2}) &&
			isAtSameSide(p1, ps, []complex128{p2, p3}) &&
			isAtSameSide(p2, ps, []complex128{p3, p1})
		if (kin || sin) && !(kin && sin) {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}

func isAtSameSide(a, b complex128, line []complex128) bool {
	sa := (real(line[1]) - real(line[0]))*(imag(a)-imag(line[0]))+
		(imag(line[1])-imag(line[0]))*(real(line[0])-real(a))
	sb := (real(line[1]) - real(line[0]))*(imag(b)-imag(line[0]))+
		(imag(line[1])-imag(line[0]))*(real(line[0])-real(b))
	return (sa > 0 && sb > 0) || (sa < 0 && sb < 0)
}
