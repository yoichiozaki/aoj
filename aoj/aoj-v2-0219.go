package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		a := make([]int, 10)
		for i := 0; i < n; i++ {
			scanner.Scan()
			c, _ := strconv.Atoi(scanner.Text())
			a[c]++
		}
		for _, v := range a {
			if v == 0 {
				fmt.Println("-")
				continue
			}
			for i := 0; i < v; i++ {
				if i == v-1 {
					fmt.Println("*")
				} else {
					fmt.Print("*")
				}
			}
		}
	}
}
