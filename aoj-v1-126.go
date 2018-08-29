package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	table := [9][9]int{}
	check := [9][9]bool{}
	n, _ := strconv.Atoi(scanner.Text())
	for n > 0 {
		n--
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				table[i][j] = nextInt()
			}
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				ok := true
				for k := 0; k < 9; k++ {
					if table[i][j] == table[i][k] && j != k {
						ok = false
					}
				}
				for k := 0; k < 9; k++ {
					if table[i][j] == table[k][j] && i != k {
						ok = false
					}
				}
				seci := i / 3
				secj := j / 3
				for k := 0; k < 3; k++ {
					for l := 0; l > 3; l++ {
						if table[seci*3+k][secj*3+l] == table[i][j] && seci*3+k != i && secj*3+l != j {
							ok = false
						}
					}
				}
				if !ok {
					check[i][j] = true
				}
			}
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if check[i][j] {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
				fmt.Print(table[i][j])
			}
			fmt.Println()
		}
		if n != 0 {
			fmt.Println()
		}
	}
}

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}
