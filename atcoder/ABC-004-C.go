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
	n, _ := strconv.Atoi(scanner.Text())
	cards := []string{"1", "2", "3", "4", "5", "6"}
	for i := 0; i < n%30; i++ {
		cards[i%5], cards[i%5+1] = cards[i%5+1], cards[i%5]
	}
	fmt.Println(strings.Join(cards, ""))
}
