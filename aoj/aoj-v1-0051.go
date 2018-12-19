
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	numbers := make([]int, 0)
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		for _, c := range input {
			numbers = append(numbers, int(c - '0'))
		}
		sort.Ints(numbers)
		minSlice := make([]string, 0)
		for _, i := range numbers {
			tmp := strconv.Itoa(i)
			minSlice = append(minSlice, tmp)
		}
		minValue, _ := strconv.Atoi(strings.Join(minSlice, ""))

		sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		maxSlice := make([]string, 0)
		for _, i := range numbers {
			tmp := strconv.Itoa(i)
			maxSlice = append(maxSlice, tmp)
		}
		maxValue, _ := strconv.Atoi(strings.Join(maxSlice, ""))
		fmt.Println(maxValue - minValue)
		numbers = []int{}
	}
}
