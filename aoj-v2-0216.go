package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	LAST_MONTH = 4280
	BASE       = 1150
)

func main() {
	for scanner.Scan() {
		w, _ := strconv.Atoi(scanner.Text())
		if w < 0 {
			break
		}
		s := BASE
		if 10 < w {
			if w <= 20 {
				s += (w - 10) * 125
			} else {
				s += 1250
			}
		}
		if 20 < w {
			if w <= 30 {
				s += (w - 20) * 140
			} else {
				s += 1400
			}
		}
		if 30 < w {
			s += (w - 30) * 160
		}
		fmt.Println(LAST_MONTH - s)
	}
}
