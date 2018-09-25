package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type bank struct {
	ID     int
	Result float64
}

type ByResult []bank

func (a ByResult) Len() int           { return len(a) }
func (a ByResult) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByResult) Less(i, j int) bool { return a[i].Result > a[j].Result }

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		banks := make(ByResult, 0)
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		for i := 1; i <= n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			b, _ := strconv.Atoi(buf[0])
			r, _ := strconv.ParseFloat(buf[1], 64)
			t, _ := strconv.Atoi(buf[2])
			bank := bank{b, 0}
			switch t {
			case 1: // 単利
				bank.Result = 10000 * (1 + float64(y)*r/100.0)
			case 2: // 複利
				bank.Result = 10000 * math.Pow(1+r/100.0, float64(y))
			}
			banks = append(banks, bank)
		}
		sort.Sort(banks)
		fmt.Println(banks[0].ID)
	}
}
