package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	flag := true
	for {
		for i := 0; i < 3; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), "")
			if buf[0] == "0" {
				flag = false
				break
			}
			a, b := 0, 0
			buf = buf[1:]
			for _, v := range buf {
				if string(v) == "A" {
					a++
				} else {
					b++
				}
			}
			if a > b {
				a++
			} else {
				b++
			}
			fmt.Println(a, b)
		}
		if !flag {
			break
		}
	}
}
