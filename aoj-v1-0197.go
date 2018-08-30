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
		buf := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(buf[0])
		b, _ := strconv.Atoi(buf[1])
		if a == 0 && b == 0 {
			break
		}
		if a < b {
			a, b = b, a
		}
		cnt := 0
		for b != 0 {
			a = a % b
			a, b = b, a
			cnt++
		}
		fmt.Println(a, cnt)
	}
}
