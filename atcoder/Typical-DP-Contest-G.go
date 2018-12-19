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
	scanner.Scan()
	s := scanner.Text()
	n := len(scanner.Text())
	dp := make([][]int64, 26)
	for i := range dp {
		dp[i] = make([]int64, n+1)
	}
	scanner.Scan()
	k, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if rune(s[i]) == 'a'+rune(j) {
				for l := 0; l < 26; l++ {
					dp[j][i] += dp[l][i+1]
					if dp[j][i] > k {
						dp[j][i] = k
						break
					}
				}
				dp[j][i]++
			} else {
				dp[j][i] = dp[j][i+1]
			}
		}
	}
	k--
	ans := make([]string, 0)
	escape := false
	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			if k < dp[j][i] {
				ans = append(ans, string('a'+j))
				for rune(s[i]) != 'a'+rune(j) {
					i++
				}
				if k == 0 {
					escape = true
				} else {
					k--
				}
				break
			} else {
				k -= dp[j][i]
			}
			if j == 25 {
				fmt.Println("Eel")
				return
			}
		}
		if escape {
			break
		}
	}
	fmt.Println(strings.Join(ans, ""))
}
