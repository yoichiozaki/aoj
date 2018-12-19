
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

func main() {
	ok := 0
	yes := true
	no := true
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	m := nextInt()
	b := make([]int, m)
	for i := range b {
		b[i] = nextInt()
	}
	c := n
	if m < c {
		c = m
	}
	for i := 0; i < c; i++ {
		if b[i] > a[i] {
			ok = 1
			break
		}
		if a[i] > b[i] {
			no = false
			break
		}
		if a[i] != b[i] {
			yes = false
		}
	}
	if yes && no {
		if m > n {
			ok = 1
		}
	}
	fmt.Println(ok)
}
