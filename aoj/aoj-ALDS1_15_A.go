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
	ret := 0
	ret += n/25
	n = n%25
	ret += n/10
	n = n%10
	ret += n/5
	n = n%5
	ret += n/1
	n = n%1
	fmt.Println(ret)
}
