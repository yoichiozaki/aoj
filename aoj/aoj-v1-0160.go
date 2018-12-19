package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	tbl := []int{600, 800, 1000, 1200, 1400, 1600}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			 break
		}
		fee := 0
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			x, _ := strconv.Atoi(buf[0])
			y, _ := strconv.Atoi(buf[1])
			h, _ := strconv.Atoi(buf[2])
			w, _ := strconv.Atoi(buf[3])
			s := x + y + h
			if s <= 160 && w <= 25 {
				k1, k2 := 0, 0
				if s <= 60 {
					k1 = 0
				} else {
					k1 = (s-61)/20 + 1
				}
				if w <= 2 {
					k2 = 0
				} else {
					k2 = (w-1)/5 + 1
				}
				if k1 < k2 {
					k1 = k2
				}
				fee += tbl[k1]
			}
		}
		fmt.Println(fee)
	}
}
