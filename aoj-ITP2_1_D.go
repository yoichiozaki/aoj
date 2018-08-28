
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
		switch nextInt() {
		case 0:
			t := nextInt()
			x := nextInt()
			A_slices[t] = append(A_slices[t], x)
			// fmt.Println(A_slices)
		case 1:
			t := nextInt()
			if len(A_slices[t]) == 0 {
				fmt.Println()
				continue
			}
			for i := 0; i < len(A_slices[t]); i++ {
				if i == len(A_slices[t])-1 {
					fmt.Printf("%d\n", A_slices[t][i])
				} else {
					fmt.Printf("%d ", A_slices[t][i])
				}
			}
		case 2:
			t := nextInt()
			if len(A_slices[t]) == 0 {
				continue
			}
			A_slices[t] = make([]int, 0)
		default:
			panic("no such operation!")
		}
	}
}
