
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
func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	for i := 0; i < n; i++ {
		a := nextInt()
		b := nextInt()
		c := nextInt()
		if a*a == b*b + c*c || b*b == a*a + c*c || c*c == a*a + b*b {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		continue
	}
}
