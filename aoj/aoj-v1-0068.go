
import (
	"bufio"
	"os"
	"strconv"
	"sort"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			os.Exit(0)
		}
		complexes := make([]complex128, 0)
		for i := 0; i < n; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), ",")
			x, _ := strconv.ParseFloat(input[0], 64)
			y, _ := strconv.ParseFloat(input[1], 64)
			complexes = append(complexes, complex(x, y))
		}
		fmt.Println(n - len(convex_hell(n, complexes)))
	}
}

func convex_hell(n int, complexes []complex128) []complex128 {
	onX := func(c1, c2 *complex128) bool {
		return real(*c1) < real(*c2)
	}
	onY := func(c1, c2 *complex128) bool {
		return imag(*c1) < imag(*c2)
	}
	OrderedBy(onY, onX).Sort(complexes)
	l := len(complexes)
	ans := make([]complex128, l+1)
	for i := range ans {
		ans[i] = 0
	}
	j := 0
	for i := 0; i < l; i++ {
		for j > 1 && cross(ans[j-1]-ans[j-2], complexes[i]-ans[j-1]) < 0 {
			j -= 1
		}
		ans[j] = complexes[i]
		j++
	}
	k := j
	for i := n-2; i > -1; i-- {
		for j > k && cross(ans[j-1]-ans[j-2], complexes[i]-ans[j-1]) < 0 {
			j -= 1
		}
		ans[j] = complexes[i]
		j += 1
	}
	return ans[0:j-1]
}

func cross(a, b complex128) float64 {
	return real(a) * imag(b) - imag(a) * real(b)
}

type lessFunc func(p1, p2 *complex128) bool

type multiSorter struct {
	complexes []complex128
	less []lessFunc
}

func (ms *multiSorter) Sort(complexes []complex128) {
	ms.complexes = complexes
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.complexes)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.complexes[i], ms.complexes[j] = ms.complexes[j], ms.complexes[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.complexes[i], &ms.complexes[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}


