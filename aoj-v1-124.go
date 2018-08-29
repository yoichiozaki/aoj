package main

import (
	"bufio"
	"fmt"
	"os"
		"strconv"
	"strings"
	"sort"
	)

var scanner = bufio.NewScanner(os.Stdin)

type Country struct {
	Name  string
	Score int
}

func (c Country) String() string {
	return fmt.Sprintf("%s,%d", c.Name, c.Score)
}

type ByScore []Country

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		countries := make([]Country, n)
		for i := range countries {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			tmp := new(Country)
			tmp.Name = input[0]
			var x int
			for j := 1; j < 4; j++ {
				switch j {
				case 1:
					x, _ = strconv.Atoi(input[j])
				case 2:
					x = 0
				case 3:
					x, _ = strconv.Atoi(input[j])
					x *= 3
				}
				tmp.Score += x
			}
			countries[i] = *tmp
		}
		sort.Sort(ByScore(countries))
		for _, v := range countries {
			fmt.Println(v)
		}
		fmt.Println()
	}
}
