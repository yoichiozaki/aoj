
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
		a := -1
		b := -1
		ans := true
		// それぞれの筒の一番上にあるたまの数字に着目すればよい
		for j := 0; j < 10; j++ {
			x := nextInt()
			if x > a {
				a = x
			} else if x > b {
				b = x
			} else {
				ans = false
			}
		}
		if ans {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
