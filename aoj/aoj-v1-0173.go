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
	for i := 0; i < 9; i++ {
		scanner.Scan()
		buf := strings.Split(scanner.Text(), " ")
		name := buf[0]
		am, _ := strconv.Atoi(buf[1])
		pm, _ := strconv.Atoi(buf[2])
		visitor := am + pm
		income := am*200 + pm*300
		fmt.Printf("%s %d %d\n", name, visitor, income)
	}
}
