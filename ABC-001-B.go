package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	if m < 100 {
		fmt.Println("00")
	} else if 100 <= m && m <= 5000 {
		fmt.Printf("%02d\n", m/100)
	} else if 6000 <= m && m <= 30000 {
		fmt.Printf("%02d\n", m/1000+50)
	} else if 35000 <= m && m <= 70000 {
		fmt.Printf("%02d\n", ((m/1000)-30)/5+80)
	} else {
		fmt.Println("89")
	}
}
