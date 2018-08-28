
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
	q := nextInt()
	set := make(map[int]struct{})
	for i := 0; i < q; i++ {
		switch nextInt() {
		case 0:
			x := nextInt()
			set[x] = struct{}{}
			fmt.Println(len(set))
		case 1:
			x := nextInt()
			if _, ok := set[x]; ok {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		case 2:
			x := nextInt()
			delete(set, x)
		default:
			panic("no such operation!")
		}
	}
}
