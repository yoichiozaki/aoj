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
	Id   int
	Time int
}

type ByTime []Team

func (a ByTime) Len() int {
	return len(a)
}
func (a ByTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByTime) Less(i, j int) bool {
	return a[i].Time < a[j].Time
}

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		teams := make([]Team, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			id, _ := strconv.Atoi(buf[0])
			m1, _ := strconv.Atoi(buf[1])
			s1, _ := strconv.Atoi(buf[2])
			m2, _ := strconv.Atoi(buf[3])
			s2, _ := strconv.Atoi(buf[4])
			m3, _ := strconv.Atoi(buf[5])
			s3, _ := strconv.Atoi(buf[6])
			m4, _ := strconv.Atoi(buf[7])
			s4, _ := strconv.Atoi(buf[8])
			team := Team{id, m1*60 + s1 + m2*60 + s2 + m3*60 + s3 + m4*60 + s4}
			teams[i] = team
		}
		sort.Sort(ByTime(teams))
		fmt.Println(teams[0].Id)
		fmt.Println(teams[1].Id)
		fmt.Println(teams[len(teams)-2].Id)
	}
}
