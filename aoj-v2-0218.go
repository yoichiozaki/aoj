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
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			math, _ := strconv.Atoi(buf[0])
			english, _ := strconv.Atoi(buf[1])
			japanese, _ := strconv.Atoi(buf[2])
			if math == 100 || english == 100 || japanese == 100 {
				fmt.Println("A")
			} else if math+english >= 180 {
				fmt.Println("A")
			} else if math+english+japanese >= 240 {
				fmt.Println("A")
			} else if math+english+japanese >= 210 {
				fmt.Println("B")
			} else if math+english+japanese >= 150 && (math >= 80 || english >= 80) {
				fmt.Println("B")
			} else {
				fmt.Println("C")
			}
		}
	}
}
