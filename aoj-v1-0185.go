package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	prime := make([]int, 1000001)
	prime[0], prime[1] = -1, -1
	for i := 2; i < 1000001; i++ {
		if prime[i] != -1 {
			prime[i] = 1 // i means "prime"
			for j := i + i; j < 1000001; j += i {
				prime[j] = -1 // -1 means "not prime"
			}
		}
	}
	for scanner.Scan() {
		ret := 0
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		for i := 1; i <= a/2; i++ {
			if prime[i] == 1 && prime[a-i] == 1 {
				ret++
			}
		}
		fmt.Println(ret)
	}
}
