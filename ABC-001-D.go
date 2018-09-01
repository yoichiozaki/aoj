package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Rain struct {
	Start int
	End   int
}
type ByStart []Rain

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }
func (r Rain) String() string        { return fmt.Sprintf("%04d-%04d", r.Start, r.End) }

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	x := make([]Rain, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		buf := strings.Split(scanner.Text(), "-")
		s, _ := strconv.Atoi(buf[0])
		e, _ := strconv.Atoi(buf[1])
		x[i].Start = s
		x[i].End = e
		sr := x[i].Start % 5
		x[i].Start -= sr
		er := x[i].End % 5
		if er > 0 {
			x[i].End += (5 - er)
		}
		if x[i].End%100 >= 60 {
			x[i].End += 40
		}
	}
	sort.Sort(ByStart(x))
	ans := make([]Rain, 0)
	start := 0
	end := 0
	for i := range x {
		if end == 0 {
			start = x[i].Start
			end = x[i].End
			continue
		}
		if x[i].Start > end {
			ans = append(ans, Rain{start, end})
			start = x[i].Start
			end = x[i].End
		}
		end = max(end, x[i].End)
	}
	ans = append(ans, Rain{start, end})
	for i := range ans {
		fmt.Println(ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
