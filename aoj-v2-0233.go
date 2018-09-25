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
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		ans := make([]string, 0)
		for a != 0 {
			b := a % 10
			if b < 0 {
				b += 10
			}
			ans = append(ans, strconv.Itoa(b))
			a -= b
			a /= -10
		}
		ans = reverse(ans)
		fmt.Println(strings.Join(ans, ""))
	}
}

func reverse(a []string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
