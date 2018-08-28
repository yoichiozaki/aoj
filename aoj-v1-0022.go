
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

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	sum := make([]int, 5001)
	for {
		n := nextInt()
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			a := nextInt()
			sum[i+1] = sum[i]+a
		}
		ret := sum[1]
		for i := 0; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				ret = max(ret, sum[j]-sum[i])
			}
		}
		fmt.Println(ret)
	}
}
