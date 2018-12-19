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
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		x, _ := strconv.Atoi(buf[1])
		if n == 0 && x == 0 {
			break
		}
		prime_table := make([]int, 1000001)
		prime_table[0] = -1
		prime_table[1] = -1
		for i := 2; i < 1000001; i++ {
			if prime_table[i] != -1 {
				prime_table[i] = 1
				for j := 2 * i; j < 1000001; j += i {
					prime_table[j] = -1 // -1 means "not prime"
				}
			}
		}
		dp := make([]int, x+1)
		dp[0] = 1
		menu := make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			menu[i], _ = strconv.Atoi(scanner.Text())
		}
		for i := 0; i < n; i++ {
			for j := 0; j <= x; j++ {
				if menu[i] <= j && dp[j-menu[i]] != 0 {
					// log.Println(dp)
					dp[j] = 1
				}
			}
		}
		flag := true
		for j := x; j >= 0; j-- {
			if prime_table[j] == 1 && dp[j] == 1 {
				flag = false
				fmt.Println(j)
				break
			}
		}
		if flag {
			fmt.Println("NA")
		}
	}
}
