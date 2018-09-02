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
	buf := strings.Split(scanner.Text(), " ")
	x, _ := strconv.Atoi(buf[0])
	y, _ := strconv.Atoi(buf[1])
	fmt.Println(y / x)
}
