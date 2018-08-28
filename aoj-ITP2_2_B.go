
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
	q := nextInt()
	A_slices := make([][]int, n)
	for i := 0; i < n; i++ {
		A_slices[i] = make([]int, 0)
	}
	for i := 0; i < q; i++ {
		// fmt.Println(A_slices)
		switch nextInt() {
		case 0:
			t := nextInt()
			x := nextInt()
			A_slices[t] = append(A_slices[t], x)
			// fmt.Println(A_slices)
		case 1:
			t := nextInt()
			if len(A_slices[t]) == 0 {
				continue
			}
			top := A_slices[t][0]
			fmt.Println(top)
		case 2:
			t := nextInt()
			if len(A_slices[t]) == 0 {
				continue
			}
			A_slices[t] = A_slices[t][1:]
		default:
			panic("no such operation!")
		}
	}
}
