package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		ans := make([]int, 7)
		for i := 0; i < n; i++ {
			scanner.Scan()
			a, _ := strconv.Atoi(scanner.Text())
			switch {
			case a < 10:
				ans[0]++
			case 10 <= a && a < 20:
				ans[1]++
			case 20 <= a && a < 30:
				ans[2]++
			case 30 <= a && a < 40:
				ans[3]++
			case 40 <= a && a < 50:
				ans[4]++
			case 50 <= a && a < 60:
				ans[5]++
			case 60 <= a:
				ans[6]++
			}
		}
		for i := range ans {
			fmt.Println(ans[i])
		}
	}
}
