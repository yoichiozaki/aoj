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
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	takoyaki := make([]int, n)
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	for i := range buf {
		takoyaki[i], _ = strconv.Atoi(buf[i])
	}
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	guests := make([]int, m)
	buf_ := strings.Split(scanner.Text(), " ")
	for i := range buf_ {
		guests[i], _ = strconv.Atoi(buf_[i])
	}
	if n < m { // たこ焼きが足りない
		fmt.Println("no")
		return
	}
	j := 0
	for i := range takoyaki {
		switch {
		case takoyaki[i]+t < guests[j]: // たこ焼きが冷えちゃった
			continue // 次のたこ焼きが売れるかどうか調べに戻る
		case guests[j] < takoyaki[i]: // 焼きあがる前に来ちゃった
			fmt.Println("no")
			return
		default:
			j++         // 売った
			if j == m { // 全員に売ることができた
				fmt.Println("yes")
				return
			}
		}
	}
	fmt.Println("no")
}
