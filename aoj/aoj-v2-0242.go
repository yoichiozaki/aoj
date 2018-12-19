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

type Entry struct {
	Word  string
	Count int
}
type lessFunc func(p1, p2 *Entry) bool
type multiSorter struct {
	entries []Entry
	less    []lessFunc
}

func (ms *multiSorter) Sort(entries []Entry) {
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
	count := func(c1, c2 *Entry) bool {
		return c1.Count > c2.Count
	}
	word := func(c1, c2 *Entry) bool {
		return c1.Word < c2.Word
	}

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		tmp := make(map[string]int)
		for i := 0; i < n; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			for _, v := range input {
				tmp[v]++
			}
		}
		dict := make([]Entry, 0)
		for word, count := range tmp {
			dict = append(dict, Entry{word, count})
		}
		scanner.Scan()
		target := scanner.Text()
		dict_ := make([]Entry, 0)
		for _, e := range dict {
			if string(e.Word[0]) == target {
				dict_ = append(dict_, Entry{e.Word, e.Count})
			}
		}
		if len(dict_) == 0 {
			fmt.Println("NA")
			return
		}
		OrderedBy(count, word).Sort(dict_)
		l := min(len(dict_), 5)
		for i := 0; i < l; i++ {
			if i != l-1 {
				fmt.Printf("%s ", dict_[i].Word)
			} else {
				fmt.Println(dict_[i].Word)
			}
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
