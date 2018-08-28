
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

func enunique(ints []int) []int {
	u := make([]int, 0, len(ints))
	m := make(map[int]bool)
	for _, val := range ints {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	a = enunique(a)
	for i := range a {
		if i == len(a)-1 {
			fmt.Println(a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
}
