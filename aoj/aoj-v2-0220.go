package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	two := make([]float64, 12)
	flag := make([]int, 12)
	for i := 0; i < 12; i++ {
		two[i] = math.Pow(2.0, 7.0-float64(i))
	}
	for scanner.Scan() {
		n, _ := strconv.ParseFloat(scanner.Text(), 64)
		if n < 0 {
			break
		}
		for i := 0; i < 12; i++ {
			flag[i] = 0
			if n >= two[i] {
				n -= two[i]
				flag[i] = 1
			}
		}
		if n < 10e-10 {
			i := 0
			for i < 12 {
				fmt.Printf("%d", flag[i])
				if i == 7 {
					fmt.Printf(".")
				}
				i++
			}
			fmt.Printf("\n")
		} else {
			fmt.Println("NA")
		}
	}
}
