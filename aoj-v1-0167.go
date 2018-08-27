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
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			arr[i], _ = strconv.Atoi(scanner.Text())
		}
		cnt := 0
		for i := n - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if arr[j] > arr[j+1] {
					// log.Println(arr)
					cnt++
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		// log.Println(arr)
		fmt.Println(cnt)
	}
}
