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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	if n%3 == 0 || strings.Contains(scanner.Text(), "3") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
