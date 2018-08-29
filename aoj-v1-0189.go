package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 0x7FFFFFFF

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		matrix := make([][]int, 10)
		for i := range matrix {
			matrix[i] = make([]int, 10)
		}
		for i := range matrix {
			for j := range matrix[i] {
				if i == j {
					matrix[i][j] = 0
				} else {
					matrix[i][j] = INF
				}
			}
		}
		num := 0
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			from, _ := strconv.Atoi(buf[0])
			to, _ := strconv.Atoi(buf[1])
			cost, _ := strconv.Atoi(buf[2])
			if num < from {
				num = from
			}
			if num < to {
				num = to
			}
			matrix[from][to] = cost
			matrix[to][from] = cost
		}
		num++
		for i := 0; i < num; i++ {
			for j := 0; j < num; j++ {
				for k := 0; k < num; k++ {
					if matrix[j][i]+matrix[i][k] < matrix[j][k] {
						matrix[j][k] = matrix[j][i] + matrix[i][k]
					}
				}
			}
		}
		ans, sum := 0, INF
		for i := 0; i < num; i++ {
			now := 0
			for j := 0; j < num; j++ {
				now += matrix[i][j]
			}
			if sum > now {
				sum = now
				ans = i
			}
		}
		fmt.Println(ans, sum)
	}
}
