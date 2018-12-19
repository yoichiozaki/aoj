package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 0x7FFFFFFF

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	tmp := INF
	for i := 0; i < n; i++ {
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())
		if t < tmp {
			tmp = t
		}
	}
	fmt.Println(tmp)
}
