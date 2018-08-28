
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
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

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

func swap(ints []int, i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func main() {
	w := nextInt()
	// println(w)
	nums := make([]int, w)
	for i := range nums {
		nums[i] = i+1
	}
	n := nextInt()
	// println(n)
	for i := 0; i < n; i++ {
		info := nextString()
		// println(info)
		infos := strings.Split(info, ",")
		a, _ := strconv.Atoi(infos[0])
		b, _ := strconv.Atoi(infos[1])
		swap(nums, a-1, b-1)
	}
	for _, v := range nums {
		fmt.Println(v)
	}
}
