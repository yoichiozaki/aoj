
import (
	"fmt"
	"sort"
)

func asc(nums ...int) []int{
	sort.Ints(nums)
	return  nums
}

func main() {
	var a, b, c int
	nums := make([]int, 3)
	fmt.Scan(&a, &b, &c)
	nums = asc(a, b, c)
	for i := range nums {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(nums[i])
	}
	fmt.Print("\n")
}
