package main

import (
	"bufio"
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for {
		success := scanner.Scan()
		if !success {
			break
		}
		flag := false
		input := scanner.Text()
		for i := 0; i < 10; i++ {
			strnum := strconv.Itoa(i)
			target := strings.Replace(input, "X", strnum, -1)
			tmp := strings.Split(target, "=")
			fs := token.NewFileSet()
			tv, err := types.Eval(fs, nil, token.NoPos, tmp[0])
			if err != nil {
				log.Println(err)
				return
			}
			val, ok := constant.Int64Val(tv.Value)
			if !ok {
				log.Println("Failed to get Int64Val")
				return
			}
			left, _ := strconv.Atoi(tmp[1])
			if left == int(val) {
				fmt.Println(i)
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("NA")
		}
	}
}
