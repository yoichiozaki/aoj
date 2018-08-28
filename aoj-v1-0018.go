
import (
	"bufio"
	"os"
	"strconv"
	"sort"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func main() {
	scanner.Split(bufio.ScanWords)
	nums := make([]int, 5)
	for i := 0; i < 5; i++ {
		nums[i] = nextInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, v := range nums {
		if i == len(nums)-1 {
			fmt.Printf("%d\n", v)
		} else {
			fmt.Printf("%d ", v)
		}
	}
}
