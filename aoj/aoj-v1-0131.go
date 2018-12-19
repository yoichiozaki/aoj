package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var (
	table, ans = [12][12]int{}, [12][12]int{}
)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		for j := 1; j < 11; j++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			tmp := [10]int{}
			for k := range input {
				tmp[k], _ = strconv.Atoi(input[k])
			}
			for l := 1; l < 11; l++ {
				table[j][l] = tmp[l-1]
			}
		}
		check(table, 1, ans)
	}
}

func check(table [12][12]int, i int, ans [12][12]int) {
	if i == 11 {
		solve(table, 2, 1, ans)
		return
	}
	ans[1][i] = 0
	check(table, i+1, ans)
	ans[1][i] = 1
	attack(table, 1, i)
	check(table, i+1, ans)
	attack(table, 1, i)
}

func solve(table [12][12]int, i, j int, ans [12][12]int) {
	if i == 11 {
		flag := true
		for k := 1; k < 11; k++ {
			if table[10][k] == 1 {
				flag = false
				break
			}
		}
		if flag {
			printans(ans)
		}
		return
	}
	if table[i-1][j] == 1 {
		ans[i][j] = 1
		attack(table, i, j)
		if j == 10 {
			solve(table, i+1, 1, ans)
		} else {
			solve(table, i, j+1, ans)
		}
		attack(table, i, j)
		ans[i][j] = 0
	} else {
		ans[i][j] = 0
		if j == 10 {
			solve(table, i+1, 1, ans)
		} else {
			solve(table, i, j+1, ans)
		}
	}
}

func printans(ans [12][12]int) {
	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			fmt.Print(ans[i][j])
			if j < 10 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func attack(table [12][12]int, i, j int) {
	table[i][j] = 1 - table[i][j]
	table[i-1][j] = 1 - table[i-1][j]
	table[i+1][j] = 1 - table[i+1][j]
	table[i][j-1] = 1 - table[i][j-1]
	table[i][j+1] = 1 - table[i][j+1]
}
