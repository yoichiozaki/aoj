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

type Team struct {
	Name                   string
	Win, Lose, Draw, Index int
}
type lessFunc func(p1, p2 *Team) bool
type multiSorter struct {
	teams []Team
	less  []lessFunc
}

func (ms *multiSorter) Sort(teams []Team) {
	ms.teams = teams
	sort.Sort(ms)
}
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}
func (ms *multiSorter) Len() int {
	return len(ms.teams)
}
func (ms *multiSorter) Swap(i, j int) {
	ms.teams[i], ms.teams[j] = ms.teams[j], ms.teams[i]
}
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.teams[i], &ms.teams[j]
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
	win := func(c1, c2 *Team) bool {
		return c1.Win > c2.Win
	}
	lose := func(c1, c2 *Team) bool {
		return c1.Lose < c2.Lose
	}
	index := func(c1, c2 *Team) bool {
		return c1.Index < c2.Index
	}

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		teams := make([]Team, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			name := buf[0]
			tmp := make([]int, n-1)
			for i := range tmp {
				tmp[i], _ = strconv.Atoi(buf[i+1])
			}
			cntwin, cntlose, cntdraw := 0, 0, 0
			for _, v := range tmp {
				switch v {
				case 1:
					cntlose++
				case 0:
					cntwin++
				case 2:
					cntdraw++
				}
			}
			teams[i] = Team{name, cntwin, cntlose, cntdraw, i}
		}
		OrderedBy(win, lose, index).Sort(teams)
		for _, t := range teams {
			fmt.Println(t.Name)
		}
	}
}
