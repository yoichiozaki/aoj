package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type Color struct {
	name    string
	R, G, B int64
}

func main() {
	colors := []Color{
		{"black", 0, 0, 0},
		{"blue", 0, 0, 0xff},
		{"lime", 0, 0xff, 0},
		{"aqua", 0, 0xff, 0xff},
		{"red", 0xff, 0, 0},
		{"fuchsia", 0xff, 0, 0xff},
		{"yellow", 0xff, 0xff, 0},
		{"white", 0xff, 0xff, 0xff},
	}
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), "")
		if buf[0] == "0" {
			break
		}
		r, _ := strconv.ParseInt(strings.Join(buf[1:3], ""), 16, 64)
		g, _ := strconv.ParseInt(strings.Join(buf[3:5], ""), 16, 64)
		b, _ := strconv.ParseInt(strings.Join(buf[5:], ""), 16, 64)
		vmin := int64(3 * 0xff * 0xff)
		var ans string
		for _, c := range colors {
			d := (r-c.R)*(r-c.R) + (g-c.G)*(g-c.G) + (b-c.B)*(b-c.B)
			if d < vmin {
				vmin = d
				ans = c.name
			}
		}
		fmt.Println(ans)
	}
}
