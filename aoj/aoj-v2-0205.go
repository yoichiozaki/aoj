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
		h := make([]int, 5)
		h[0], _ = strconv.Atoi(scanner.Text())
		if h[0] == 0 {
			break
		}
		for i := 1; i <= 4; i++ {
			scanner.Scan()
			h[i], _ = strconv.Atoi(scanner.Text())
		}
		tbl := make([]bool, 3)
		aiko := false
		for _, v := range h {
			tbl[v-1] = true
		}
		if tbl[0] && tbl[1] && tbl[2] { // 全ての手があるあいこ
			aiko = true
		}
		if h[0] == h[1] && h[1] == h[2] && h[2] == h[3] && h[3] == h[4] {
			aiko = true
		}
		if aiko {
			fmt.Printf("3\n3\n3\n3\n3\n")
		} else {
			win := 0
			if !tbl[0] {
				win = 2
			}
			if !tbl[1] {
				win = 3
			}
			if !tbl[2] {
				win = 1
			}
			for i := 0; i < 5; i++ {
				if h[i] == win {
					fmt.Println(1)
				} else {
					fmt.Println(2)
				}
			}
		}
	}
}
