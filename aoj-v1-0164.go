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
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		scanner.Scan()
		buf := strings.Split(scanner.Text(), " ")
		input := make([]int, len(buf))
		for i := range input {
			input[i], _ = strconv.Atoi(buf[i])
		}
		fmt.Println(31)
		i, k := 0, 31
		for k > 1 {
			k -= input[i]
			fmt.Println(k)
			k = k - (k-1)%5
			fmt.Println(k)
			i++
			if i >= n {
				i = 0
			}
		}
		fmt.Println(0)
	}
}
