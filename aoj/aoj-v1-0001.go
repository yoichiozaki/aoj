
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
	heighs := make([]int, 10)
	for i := range heighs {
		heighs[i] = nextInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(heighs)))
	for i := 0; i < 3; i++ {
		fmt.Println(heighs[i])
	}
}
