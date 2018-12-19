package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	D, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	N := scanner.Text()
	Ns := strings.Split(scanner.Text(), "")
	// log.Println(Ns)

	// dp1[i][j]: 0, 1, ..., 99...99(9がj個)の数字のうち、各桁の和をDで割った余りがiとなる数字の個数
	// k = 0, 1, 2, ..., 8, 9に対して
	// dp1[(i + k) mod D][j] += dp1[i][j-1]
	// dp1[0][0] = 1
	dp1 := make([][]int, D)
	for i := range dp1 {
		dp1[i] = make([]int, digit(N)+1)
	}
	dp1[0][0] = 1
	for j := 1; j < digit(N)+1; j++ {
		for i := 0; i < D; i++ {
			for k := 0; k < 10; k++ {
				dp1[(i+k)%D][j] += dp1[i][j-1]
			}
		}
	}

	// log.Println(dp1)

	// dp2[i][j]: Nのj桁目までの数字(例えばN = 123456789に対して3桁目までの数字とは0, 1, 2, ..., 788, 789)
	// のうち各桁の和をDで割った余りがiになる数字の個数
	// k = 0, 1, 2, ..., 8, 9に対して
	// k = Nのj桁目のとき, dp2[(i + k) mod D][j] += dp2[i][j-1]
	// k < Nのj桁目のとき, dp2[(i + k) mod D][j] += dp1[i][j-1]
	// Nのj桁目 < kのとき, dp2[(i + k) mod D][j] += 0
	// dp2[0][0] = 1
	dp2 := make([][]int, D)
	for i := range dp2 {
		dp2[i] = make([]int, digit(N)+1)
	}
	dp2[0][0] = 1

	for j := 1; j <= digit(N); j++ {
		for i := 0; i < D; i++ {
			for k := 0; k < 10; k++ {
				if k == string2int(Ns[len(Ns)-j]) {
					dp2[(i+k)%D][j] += dp2[i][j-1]
					// log.Printf("dp2[(%d + %d) mod %d][%d] += dp2[%d][%d]", i, k, D, j, i, j-1)
				} else if k < string2int(Ns[len(Ns)-j]) {
					dp2[(i+k)%D][j] += dp1[i][j-1]
					// log.Printf("dp2[(%d + %d) mod %d][%d] += dp1[%d][%d]", i, k, D, j, i, j-1)
				} else {
					dp2[(i+k)%D][j] += 0
				}
				// log.Printf("i: %d, j: %d, k: %d\n", i, j, k)
				// log.Println(dp2)
			}

			log.Printf("i: %d, j: %d\n", i, j)
			fmt.Println(dp2[0][digit(N)] - 1)
			if dp2[0][digit(N)]-1 == 468357804 {
				fmt.Println("OK")
			}

		}
	}
	// log.Println(dp2)
	// dp2[0][Nの桁数] - 1
	fmt.Println(dp2[0][digit(N)] - 1)
	if dp2[0][digit(N)]-1 == 468357804 {
		fmt.Println("OK")
	}
}

func digit(N string) int {
	return len(N)
}

func string2int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
