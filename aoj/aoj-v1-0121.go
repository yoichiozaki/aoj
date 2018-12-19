
import (
	"bufio"
	"os"
	"strings"
		"fmt"
	)

var scanner = bufio.NewScanner(os.Stdin)

var swap = [][]int{
	{1, 4},
	{0, 2, 5},
	{1, 3, 6},
	{2, 7},
	{0, 5},
	{1, 4, 6},
	{2, 5, 7},
	{3, 6},
}

type entry struct {
	arr []string
	cnt int
}

func bfs() map[string]int {
	ia := []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	count := make(map[string]int)
	count[strings.Join(ia, "")] = 0
	q := make([]entry, 0)
	q = append(q, entry{ia, 0})
	for len(q) != 0 {
		tmp := q[0]
		q = q[1:]
		a := tmp.arr
		c := tmp.cnt
		i := index(a, "0")
		for _, j := range swap[i] {
			na := make([]string, len(a), cap(a))
			copy(na, a)
			na[i], na[j] = na[j], na[i]
			if _, ok := count[strings.Join(na, "")]; !ok {
				count[strings.Join(na, "")] = c+1
				q = append(q, entry{na, c+1})
			}
		}
	}
	return count
}

func main() {
	count := bfs()
	// log.Println(count)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		fmt.Println(count[strings.Join(a, "")])
	}
}

func index(a []string, target string) int {
	for i, v := range a {
		if v == target {
			return i
		}
	}
	return -1
}
