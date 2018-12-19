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
HERE:
	for {
		max := 0
		shop := 'A'
		for i := 0; i < 5; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(buf[0])
			b, _ := strconv.Atoi(buf[1])
			if a == 0 && b == 0 {
				break HERE
			}
			if max < a+b {
				max = a + b
				shop = 'A' + int32(i)
			}
		}
		fmt.Println(string(shop), max)
	}
}
