
import (
	"fmt"
	"sort"
)

func main() {
	var n, sum int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	for _, v := range a {
		sum += v
	}
	fmt.Printf("%d %d %d\n", a[0], a[n-1], sum)
}
