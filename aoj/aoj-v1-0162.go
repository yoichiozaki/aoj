package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	MAX = 1000000
)

func main() {
	a := make([]int, MAX+1)
	a[1] = 1
	for i := 2; i <= MAX; i *= 2 {
		a[i]++
	}
	for i := 1; i <= MAX; i++ {
		if a[i] != 0 {
			for j := i * 3; j <= MAX; j *= 3 {
				if a[j] == 0 {
					a[j]++
				}
			}
		}
	}
	for i := 1; i <= MAX; i++ {
		if a[i] != 0 {
			for j := i * 5; j <= MAX; j *= 3 {
				if a[j] == 0 {
					a[j]++
				}
			}
		}
	}
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(buf[0])
		if m == 0 {
			break
		}
		n, _ := strconv.Atoi(buf[1])
		cnt := 0
		for i := m; i <= n; i++ {
			if a[i] != 0 {
				cnt++
			}
		}
		fmt.Println(cnt)
	}
}
