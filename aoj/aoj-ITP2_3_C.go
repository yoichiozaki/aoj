
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

func count(ints []int, target int) int {
	var ret int
	for _, v := range ints {
		if v == target {
			ret++
		}
	}
	return ret
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
		k := nextInt()
		fmt.Printf("%d\n", count(a[b:e], k))
	}
}
