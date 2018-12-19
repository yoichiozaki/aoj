
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

func swap(ints []int, b, e, t int) []int {
	for k := 0; k < e-b; k++ {
		ints[b+k], ints[t+k] = ints[t+k], ints[b+k]
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
		t := nextInt()
		a = swap(a, b, e, t)
	}
	for i := range a {
		if i == len(a)-1 {
			fmt.Println(a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
}
