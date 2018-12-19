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
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		cards := make([]int, len(buf))
		for i := range buf {
			cards[i], _ = strconv.Atoi(buf[i])
		}
		if cards[0] == 0 {
			break
		}
		sum, ace := 0, 0
		for _, v := range cards {
			if v == 1 {
				v = 0
				ace++
			} else if v >= 10 {
				v = 10
			}
			sum += v
			if sum+ace > 21 {
				break
			}
		}
		if sum+ace > 21 {
			fmt.Println(0)
		} else if ace > 0 && sum+ace <= 11 {
			fmt.Println(sum + ace + 10) // 1->11なるAceは高々1枚
		} else {
			fmt.Println(sum + ace)
		}
	}
}
