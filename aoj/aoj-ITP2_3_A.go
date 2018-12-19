
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

func getMax(ints ...int) int {
	var max int
	max = ints[0]
	for _, v := range ints {
		if max < v {
			max = v
		}
	}
	return max
}

func getMin(ints ...int) int {
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
	a := nextInt()
	b := nextInt()
	c := nextInt()
	fmt.Printf("%d %d\n", getMin(a, b, c), getMax(a, b, c))
}
