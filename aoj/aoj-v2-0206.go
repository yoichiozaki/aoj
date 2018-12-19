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
		val := 0
		L, _ := strconv.Atoi(scanner.Text())
		if L == 0 {
			break
		}
		for i := 0; i < 12; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			M, _ := strconv.Atoi(buf[0])
			N, _ := strconv.Atoi(buf[1])
			L -= (M - N)
			if L <= 0 && val == 0 {
				val = i + 1
			}
		}
		if val == 0 {
			fmt.Println("NA")
		} else {
			fmt.Println(val)
		}
	}
}
