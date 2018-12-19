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
		_a, _ := strconv.Atoi(buf[0])
		_b, _ := strconv.Atoi(buf[1])
		a := strings.Split(buf[0], "")
		b := strings.Split(buf[1], "")
		if _a == 0 && _b == 0 {
			break
		}
		hit := 0
		blow := 0
		for i, av := range a {
			for j, bv := range b {
				if i == j && av == bv {
					hit++
				}
				if i != j && av == bv {
					blow++
				}
			}
		}
		fmt.Println(hit, blow)
	}
}
