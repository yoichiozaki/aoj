
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

type pair struct {
	first, second int
}

var st = make([]pair, 2097152)

func main() {
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	// n, _ := strconv.Atoi(input[0])
	q, _ := strconv.Atoi(input[1])
	for i := 0; i < 1048576; i++ {
		st[1048576+i].first=i+1
	}
	for i := 1048575; i >= 1; i-- {
		st[i].first=st[i*2].first
	}
	for i := 0; i < q; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(input[0])
		v, _ := strconv.Atoi(input[1])
		a--
		add(a, v)
		fmt.Printf("%d %d\n", st[1].first,st[1].second)
	}
}

func add(a, b int) {
	a += 1048576
	st[a].second += b
	a /= 2
	for a != 0 {
		if st[a*2].second<st[a*2+1].second {
			st[a]=st[a*2+1]
		} else {
			st[a]=st[a*2]
		}
		a /= 2
	}
}
