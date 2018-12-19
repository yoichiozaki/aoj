
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type entry struct {
	id, score int
}

type lessFunc func(p1, p2 *entry) bool

type multiSorter struct {
	entries []entry
	less    []lessFunc
}

func (ms *multiSorter) sort(entries []entry) {
	ms.entries = entries
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.entries)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.entries[i], ms.entries[j] = ms.entries[j], ms.entries[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.entries[i], &ms.entries[j]
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

func main() {
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		if m == 0 {
			break
		}
		results := make([]entry, 0)
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			data := make([]int, len(buf))
			for i := range data {
				data[i], _ = strconv.Atoi(buf[i])
			}
			id, score := score(data)
			results = append(results, entry{id, score})
		}
		score_axis := func(e1, e2 *entry) bool {
			return e1.score > e2.score
		}
		id_axis := func(e1, e2 *entry) bool {
			return e1.id < e2.id
		}
		OrderedBy(score_axis, id_axis).sort(results)
		// log.Println(results)
		for _, e := range results {
			fmt.Printf("%d %d\n", e.id, e.score)
		}
	}
}

func score(input []int) (int, int) {
	id := input[0]
	data := input[1:]
	frame := 1
	scores := make([]int, 11)
	for frame <= 10 {
		if data[0] == 10 {
			scores[frame] = data[0] + data[1] + data[2]
			data = data[1:]
			frame += 1
		} else if data[0]+data[1] == 10 {
			scores[frame] = data[0] + data[1] + data[2]
			data = data[2:]
			frame += 1
		} else {
			scores[frame] = data[0] + data[1]
			data = data[2:]
			frame += 1
		}
	}
	sum := 0
	for i, v := range scores {
		if i != 0 {
			sum += v
		}
	}
	return id, sum
}

