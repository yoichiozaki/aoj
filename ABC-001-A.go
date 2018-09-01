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
	h1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	h2, _ := strconv.Atoi(scanner.Text())
	fmt.Println(h1 - h2)
}
