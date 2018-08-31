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

type Patient struct {
	ID, Distance int
}
type ByDistance []Patient

func (p ByDistance) Len() int           { return len(p) }
func (p ByDistance) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByDistance) Less(i, j int) bool { return p[i].Distance > p[j].Distance }

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		patients := make([]Patient, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			id, _ := strconv.Atoi(buf[0])
			am, _ := strconv.Atoi(buf[1])
			pm, _ := strconv.Atoi(buf[2])
			patients[i] = Patient{id, am + pm}
		}
		sort.Sort(ByDistance(patients))
		fmt.Println(patients[0].ID, patients[0].Distance)
	}
}
