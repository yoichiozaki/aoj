
import (
	"fmt"
	"sort"
)

type num []int
type nums [][]int

func (n nums) Len() int{
	return len(n)
}

func (n nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n nums) Less(i, j int) bool {
	ni := n[i][1]
	nj := n[j][1]
	return ni > nj
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make(nums, n)
	for i := range a {
		a[i] = make(num, 2)
		a[i][1] = i
		fmt.Scan(&a[i][0])
	}
	// fmt.Println(a)
	sort.Sort(a)
	// fmt.Println(a)

	for i, v := range a {
		if i == 0 {
			fmt.Printf("%d", v[0])
		} else {
			fmt.Printf(" %d", v[0])
		}
		if i == n - 1 {
			fmt.Print("\n")
		}
	}
}
