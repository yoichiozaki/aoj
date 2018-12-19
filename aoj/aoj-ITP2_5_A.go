
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

type Point struct {
	x int
	y int
}

type Points []Point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	if p[i].x == p[j].x {
		return p[i].y < p[j].y
	}
	return p[i].x < p[j].x
}

func main() {
	scanner.Split(bufio.ScanWords)
	var points Points = []Point{}
	n := nextInt()
	for i := 0; i < n; i++ {
		x := nextInt()
		y := nextInt()
		points = append(points, Point{x, y})
	}
	sort.Sort(points)
	for i := range points {
		fmt.Printf("%d %d\n", points[i].x, points[i].y)
	}
}
