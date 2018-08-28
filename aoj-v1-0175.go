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
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == -1 {
			break
		}
		if n == 0 {
			fmt.Println(0)
			continue
		}
		ans := make([]int, 0)
		for n > 0 {
			ans = append([]int{n % 4}, ans...)
			n = n / 4
		}

		buf := make([]string, len(ans))
		for i := range ans {
			buf[i] = strconv.Itoa(ans[i])
		}
		fmt.Println(strings.Join(buf, ""))
	}
}
