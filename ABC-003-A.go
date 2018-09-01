package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	ans := 0
	for i := 1; i <= n; i++ {
		ans += i * 10000
	}
	fmt.Println(ans / n)
}
