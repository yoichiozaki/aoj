package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	xa, _ := strconv.ParseFloat(buf[0], 64)
	ya, _ := strconv.ParseFloat(buf[1], 64)
	xb, _ := strconv.ParseFloat(buf[2], 64)
	yb, _ := strconv.ParseFloat(buf[3], 64)
	xc, _ := strconv.ParseFloat(buf[4], 64)
	yc, _ := strconv.ParseFloat(buf[5], 64)
	fmt.Printf("%.1f", math.Abs(((xb-xa)*(yc-ya)-(yb-ya)*(xc-xa))/2))
}
