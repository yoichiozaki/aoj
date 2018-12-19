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

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		m, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		scanner.Scan()
		buf_ := strings.Split(scanner.Text(), " ")
		items := make([]int, n)
		for i := 0; i < n; i++ {
			items[i], _ = strconv.Atoi(buf_[i])
		}
		sort.Ints(items)
		sum := 0
		for _, v := range items {
			sum += v
		}
		for i := n - m; i >= 0; i -= m {
			sum -= items[i]
		}
		fmt.Println(sum)
	}
}
