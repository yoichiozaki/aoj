package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	table := []int{0, 1, 2, 3, 5, 7, 8, 9}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		out := make([]int, 10)
		now := 0
		for n > 0 {
			out[now] = table[n%8]
			now++
			n /= 8
		}
		tmp := make([]string, 0)
		for i := now - 1; i >= 0; i-- {
			tmp = append(tmp, strconv.Itoa(out[i]))
		}
		fmt.Println(strings.Join(tmp, ""))
	}
}
