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
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		s, _ := strconv.Atoi(input[0])
		e, _ := strconv.Atoi(input[1])
		a := make([]int, 0)
		if s <= 5 {
			if s > e {
				for j := s; j > e-1; j-- {
					a = append(a, j)
				}
			} else {
				for j := s; j < e+1; j++ {
					a = append(a, j)
				}
			}
		} else if s < e {
			for j := s; j < 10; j++ {
				a = append(a, j)
			}
		} else {
			for j := s; j < 10; j++ {
				a = append(a, j)
			}
			if e <= 5 {
				for j := 5; j > e-1; j-- {
					a = append(a, j)
				}
			} else {
				for j := 5; j > 0; j-- {
					a = append(a, j)
				}
				for j := 0; j < e+1; j++ {
					a = append(a, j)
				}
			}
		}
		for i, v := range a {
			if i != len(a)-1 {
				fmt.Printf("%d ", v)
			} else {
				fmt.Println(v)
			}
		}
	}
}
