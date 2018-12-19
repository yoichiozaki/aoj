
import (
	"bufio"
	"os"
	"strconv"
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

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n*factorial(n-1)
}
func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	fmt.Println(factorial(n))
}
