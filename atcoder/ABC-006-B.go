package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
const DIV = 10007

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	// fmt.Println(toribocchi(n))
	a := make([]int, n)
	switch n {
	case 1, 2:
		fmt.Println(0)
		return
	case 3:
		fmt.Println(1)
		return
	default:
		a[0], a[1], a[2] = 0, 0, 1
		for i := 3; i < n; i++ {
			a[i] = (a[i-1] + a[i-2] + a[i-3])%DIV
		}
		fmt.Println(a[n-1]%DIV)
	}
}

// func toribocchi(n int) int {
// 	switch n {
// 	case 1:
// 		return 0
// 	case 2:
// 		return 0
// 	case 3:
// 		return 1
// 	default:
// 		return (toribocchi(n-1) + toribocchi(n-2) + toribocchi(n-3))%DIV
// 	}
// }
