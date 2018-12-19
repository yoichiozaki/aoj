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
		c, _ := strconv.Atoi(buf[2])
		d, _ := strconv.Atoi(buf[3])
		e, _ := strconv.Atoi(buf[4])
		f, _ := strconv.Atoi(buf[5])
		if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 && f == 0 {
			break
		}
		fmt.Println(100 - ((5*a + 3*b) * 2) - (f-5*a-3*b-e)*3 + 15*a + 15*b + 15*(5*a+3*b) + c*7 + d*2)
	}
}
