
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

func swap(ints []int, b int, e int) []int {
	l := e-b
	for i := 0; i < l/2; i++ {
		ints[b+i], ints[e-1-i] = ints[e-1-i], ints[b+i]
	}
	return ints
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
		b := nextInt()
		e := nextInt()
		a = swap(a, b, e)
	}
	for i := range a {
		if i == len(a)-1 {
			fmt.Println(a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
}
