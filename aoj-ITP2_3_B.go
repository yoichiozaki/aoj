
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)


func nextInt() int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func getMax(ints []int) int {
	var max int
	max = ints[0]
	for _, v := range ints {
		if max < v {
			max = v
		}
	}
	return max
}

func getMin(ints []int) int {
	var min int
	min = ints[0]
	for _, v := range ints {
		if min > v {
			min = v
		}
	}
	return min
}
func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	q := nextInt()
	for i := 0; i < q; i++ {
		switch nextInt() {
		case 0:
			// min
			b := nextInt()
			e := nextInt()
			fmt.Printf("%d\n", getMin(a[b:e]))
		case 1:
			// max
			b := nextInt()
			e := nextInt()
			fmt.Printf("%d\n", getMax(a[b:e]))

		default:
			panic("no such operation!")
		}
	}
}
