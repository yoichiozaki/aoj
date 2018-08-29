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
	table := [6][5]string{
		{"a", "b", "c", "d", "e"},
		{"f", "g", "h", "i", "j"},
		{"k", "l", "m", "n", "o"},
		{"p", "q", "r", "s", "t"},
		{"u", "v", "w", "x", "y"},
		{"z", ".", "?", "!", " "},
	}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")
		if len(input)%2 != 0 {
			fmt.Println("NA")
			continue
		}
		ret := make([]string, 0)
		for i := 0; i < len(input); i += 2 {
			row, _ := strconv.Atoi(input[i])
			col, _ := strconv.Atoi(input[i+1])
			ret = append(ret, table[row-1][col-1])
		}
		fmt.Println(strings.Join(ret, ""))
	}
}
