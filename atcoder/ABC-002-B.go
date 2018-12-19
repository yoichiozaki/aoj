package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	input := scanner.Text()
	ans := make([]string, 0)
	vowel := make(map[string]struct{})
	vowel["a"] = struct{}{}
	vowel["i"] = struct{}{}
	vowel["u"] = struct{}{}
	vowel["e"] = struct{}{}
	vowel["o"] = struct{}{}
	for _, v := range input {
		if _, ok := vowel[string(v)]; !ok {
			ans = append(ans, string(v))
		}
	}
	fmt.Println(strings.Join(ans, ""))
}
