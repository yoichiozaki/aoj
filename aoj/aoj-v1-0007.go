
import (
	"bufio"
	"os"
	"strconv"
	"math"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func calc(n int) float64 {
	tmp := 100000.0
	for i := 0; i < n; i++ {
		tmp *= 1.05
		tmp = ceil(tmp, -3)
	}
	return tmp
}

func ceil(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Ceil(f*shift)/shift
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	ans := calc(n)
	fmt.Println(int(ans))
}
