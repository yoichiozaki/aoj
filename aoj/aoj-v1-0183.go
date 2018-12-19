package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	allb = "bbb"
	allw = "www"
)

func main() {
	for scanner.Scan() {
		row1 := scanner.Text()
		if row1 == "0" {
			break
		}
		scanner.Scan()
		row2 := scanner.Text()
		scanner.Scan()
		row3 := scanner.Text()
		x := make([]string, 8)
		x[0] = row1
		x[1] = row2
		x[2] = row3
		for i := 0; i < 3; i++ {
			col := make([]string, 3)
			col[0] = string(row1[i])
			col[1] = string(row2[i])
			col[2] = string(row3[i])
			x[i+3] = strings.Join(col, "")
		}
		tmp := make([]string, 3)
		tmp[0] = string(row1[0])
		tmp[1] = string(row2[1])
		tmp[2] = string(row3[2])
		x[6] = strings.Join(tmp, "")
		tmp[0] = string(row1[2])
		tmp[1] = string(row2[1])
		tmp[2] = string(row3[0])
		x[7] = strings.Join(tmp, "")
		flag := false
		for _, v := range x {
			if string(v) == allb {
				fmt.Println("b")
				flag = true
				break
			} else if string(v) == allw {
				fmt.Println("w")
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("NA")
		}
	}
}
