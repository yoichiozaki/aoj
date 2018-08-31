package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	table := [][]int{ // abcdefgの順番の
		{1, 1, 1, 1, 1, 1, 0}, // 1
		{0, 1, 1, 0, 0, 0, 0}, // 2
		{1, 1, 0, 1, 1, 0, 1}, // 3
		{1, 1, 1, 1, 0, 0, 1}, // 4
		{0, 1, 1, 0, 0, 1, 1}, // 5
		{1, 0, 1, 1, 0, 1, 1}, // 6
		{1, 0, 1, 1, 1, 1, 1}, // 7
		{1, 1, 1, 0, 0, 1, 0}, // 8
		{1, 1, 1, 1, 1, 1, 1}, // 9
		{1, 1, 1, 1, 0, 1, 1}, // 10
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == -1 {
			break
		}
		initial := []int{0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < n; i++ {
			scanner.Scan()
			d, _ := strconv.Atoi(scanner.Text())
			for j := 6; j >= 0; j-- {
				fmt.Printf("%d", initial[j]^table[d][j])
				initial[j] = table[d][j]
			}
			fmt.Println()
		}
	}
}
