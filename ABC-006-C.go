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
	input := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(input[0])
	m, _ := strconv.Atoi(input[1])
	baby, man, old := 0, 0, 0
	q := m / n // 一個体あたりの平均足本数
	r := m % n
	if q < 2 || 4 < q || (q == 4 && 0 < r) {
		baby, man, old = -1, -1 ,-1
	} else {
		if q == 2 {
			man = n - r
			old = r
			baby = 0
		} else if q == 3 {
			man = 0
			old = n - r
			baby = r
		} else {
			man, old = 0, 0
			baby = n
		}
	}
	fmt.Println(man, old, baby)
}