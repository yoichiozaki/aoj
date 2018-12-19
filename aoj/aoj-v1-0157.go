
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sort"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		hr := make([]entry, 0)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			first, _ := strconv.Atoi(buf[0])
			second, _ := strconv.Atoi(buf[1])
			mt := entry{first, second}
			hr = append(hr, mt)
		}
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			first, _ := strconv.Atoi(buf[0])
			second, _ := strconv.Atoi(buf[1])
			mt := entry{first, second}
			hr = append(hr, mt)
		}
		sort.Sort(ByFirst(hr))
		fmt.Println(LIS(hr))
	}
}

type entry struct {
	first, second int
}

type ByFirst []entry

func(f ByFirst) Len() int {
	return len(f)
}

func (f ByFirst) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f ByFirst) Less(i, j int) bool {
	return f[i].first < f[j].first
}

func LIS(hr []entry) int {
	n := len(hr)
	lis := make([]int, n)
	lis[n-1] = 1
	for i := n-2; i > -1; i-- {
		m := 0
		for j := i+1; j < n; j++ {
			if hr[i].first < hr[j].first && hr[i].second < hr[j].second && lis[j] > m {
				m = lis[j]
			}
		}
		lis[i] = m+1
	}
	return lis[0]
}
