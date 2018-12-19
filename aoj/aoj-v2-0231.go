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

type Pair struct {
	first, second int
}
type ByFirst []Pair

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	for {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		man := make(ByFirst, 2*a)
		for i := 0; i < a; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			b, _ := strconv.Atoi(buf[0])
			c, _ := strconv.Atoi(buf[1])
			d, _ := strconv.Atoi(buf[2])
			man[2*i] = Pair{c, b}
			man[2*i+1] = Pair{d, -b}
		}
		sort.Sort(man)
		ok := true
		now := 0
		for i := 0; i < a*2; i++ {
			now += man[i].second
			if now > 150 {
				ok = false
			}
		}
		if ok {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}
