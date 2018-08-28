
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

type Item struct {
	v int
	w int
	t string
	d int
	s string
}

type Items []Item

func (p Items) Len() int {
	return len(p)
}

func (p Items) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Items) Less(i, j int) bool {
	if p[i].v == p[j].v {
		if p[i].w == p[j].w {
			if p[i].t == p[j].t {
				if p[i].d == p[j].d {
					return p[i].s < p[j].s
				} else {
					return p[i].d < p[j].d
				}
			} else {
				return p[i].t < p[j].t
			}
		} else {
			return p[i].w < p[j].w
		}
	}
	return p[i].v < p[j].v
}

func main() {
	scanner.Split(bufio.ScanWords)
	var points Items = []Item{}
	n := nextInt()
	for i := 0; i < n; i++ {
		v := nextInt()
		w := nextInt()
		t := nextString()
		d := nextInt()
		s := nextString()
		points = append(points, Item{v,w,t,d,s})
	}
	sort.Sort(points)
	for i := range points {
		fmt.Printf("%d %d %s %d %s\n", points[i].v, points[i].w, points[i].t, points[i].d, points[i].s)
	}
}
