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
		having := map[string]Entry{}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			name := buf[0]
			price, _ := strconv.Atoi(buf[1])
			having[name] = Entry{price, []string{}}
		}
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			name_, price_ := buf_[0], INF
			if _, ok := having[name_]; ok {
				price_ = having[name_].price
			}
			having[name_] = Entry{price_, buf_[2:]}
		}
		scanner.Scan()
		fmt.Println(calculation(scanner.Text(), having))
	}
}

type Entry struct {
	price    int
	elements []string
}

// メモ再帰的
func calculation(target string, having map[string]Entry) int {
	ans := INF
	if len(having[target].elements) > 0 {
		ans = 0
		for _, item := range having[target].elements {
			ans += calculation(item, having)
		}
	}
	return min(ans, having[target].price)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
