package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

const SIZE = 10000000

func main() {
	prime := make([]bool, SIZE+1)
	ans := make([]int, SIZE+1)
	for i := 0; i < SIZE+1; i++ {
		prime[i] = true
	}
	prime[0], prime[1] = false, false
	for i := 2; i*i < SIZE+1; i++ {
		if !prime[i] {
			continue
		}
		for j := i * i; j < SIZE+1; j += i {
			prime[j] = false
		}
	}
	ans[13] = 13
	for i := 14; i < SIZE+1; i++ {
		ans[i] = ans[i-1]
		if prime[i] && prime[i-2] && prime[i-6] && prime[i-8] {
			ans[i] = i
		}
	}
	for {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		fmt.Println(ans[n])
	}
}
