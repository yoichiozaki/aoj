// binary search

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
		a := make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			a[i], _ = strconv.Atoi(scanner.Text())
		}
		scanner.Scan()
		k, _ := strconv.Atoi(scanner.Text())
		cnt := 0
		low, high := 0, n-1
		for low <= high {
			cnt++
			mid := (low + high) >> 1
			if k == a[mid] {
				break
			}
			if k < a[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		fmt.Println(cnt)
	}
}
