package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

var (
	str [33]rune
	UF  [26]int
	in  [26]int
	out [26]int
)

func find(a int) int {
	if UF[a] < 0 {
		return a
	}
	UF[a] = find(UF[a])
	return UF[a]
}

func union(a, b int) {
	a = find(a)
	b = find(b)
	if a == b {
		return
	}
	UF[a] += UF[b]
	UF[b] = a
}

func main() {
	for {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		for i := 0; i < 26; i++ {
			in[i], out[i] = 0, 0
			UF[i] = -1
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			str := scanner.Text()
			in[str[0]-'a']++
			out[str[len(str)-1]-'a']++
			union(int(str[0]-'a'), int(str[len(str)-1]-'a'))
		}
		flag := false
		parent := -1
		for i := 0; i < 26; i++ {
			if in[i] != out[i] {
				flag = true
			} else if in[i] > 0 {
				if parent == -1 {
					parent = find(i)
				} else if parent != find(i) {
					flag = true
				}
			}
		}
		if flag {
			fmt.Println("NG")
		} else {
			fmt.Println("OK")
		}
	}
}
