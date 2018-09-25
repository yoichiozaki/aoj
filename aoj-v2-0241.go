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
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			x1, _ := strconv.Atoi(buf[0])
			y1, _ := strconv.Atoi(buf[1])
			z1, _ := strconv.Atoi(buf[2])
			w1, _ := strconv.Atoi(buf[3])
			x2, _ := strconv.Atoi(buf[4])
			y2, _ := strconv.Atoi(buf[5])
			z2, _ := strconv.Atoi(buf[6])
			w2, _ := strconv.Atoi(buf[7])
			x3 := x1*x2 - y1*y2 - z1*z2 - w1*w2
			y3 := x1*y2 + y1*x2 + z1*w2 - w1*z2
			z3 := x1*z2 - y1*w2 + z1*x2 + w1*y2
			w3 := x1*w2 + y1*z2 - z1*y2 + w1*x2
			fmt.Println(x3, y3, z3, w3)
		}
	}
}
