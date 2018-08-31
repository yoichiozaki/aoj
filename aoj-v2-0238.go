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
		t, _ := strconv.Atoi(scanner.Text())
		if t == 0 {
			break
		}
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			s, _ := strconv.Atoi(buf[0])
			f, _ := strconv.Atoi(buf[1])
			t -= (f - s)
		}
		if t <= 0 {
			fmt.Println("OK")
		} else {
			fmt.Println(t)
		}
	}
}
