package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	t := scanner.Text()
	if s == t && !hasAt(s) && !hasAt(t) { // 同じ文字列かつ@を含んでいない
		fmt.Println("You can win")
		return
	}
	if s != t && !hasAt(s) && !hasAt(t) { // 異なる文字列だが共に@を含まない
		fmt.Println("You will lose")
		return
	}
	for i, _ := range s {
		si := s[i]
		ti := t[i]
		if si != ti {
			if si != '@' && ti != '@' { // 互いに異なる文字が@ではない
				fmt.Println("You will lose")
				return
			} else if si == '@' { // 互いに異なる文字が@
				if !check(ti) { // @がatcoderのいずれかの文字で埋まれば良い
					fmt.Println("You will lose")
					return
				}
			} else {
				if !check(si) {
					fmt.Println("You will lose")
					return
				}
			}
		}
	}
	fmt.Println("You can win")
}

func hasAt(x string) bool {
	for i, _ := range x {
		if x[i] == '@' {
			return true
		}
	}
	return false
}

func check(b byte) bool {
	if b == 'a' || b == 't' || b == 'c' || b == 'o' || b == 'd' || b == 'e' || b == 'r' {
		return true
	}
	return false
}
