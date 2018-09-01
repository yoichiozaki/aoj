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
	deg, _ := strconv.ParseFloat(buf[0], 64)
	dis, _ := strconv.ParseFloat(buf[1], 64)
	var dir string
	var w int
	deg /= 10
	dis /= 60
	switch {
	case deg >= 11.25 && deg < 33.75:
		dir = "NNE"
	case deg >= 33.75 && deg < 56.25:
		dir = "NE"
	case deg >= 56.25 && deg < 78.75:
		dir = "ENE"
	case deg >= 78.75 && deg < 101.25:
		dir = "E"
	case deg >= 101.25 && deg < 123.75:
		dir = "ESE"
	case deg >= 123.75 && deg < 146.25:
		dir = "SE"
	case deg >= 146.25 && deg < 168.75:
		dir = "SSE"
	case deg >= 168.75 && deg < 191.25:
		dir = "S"
	case deg >= 191.25 && deg < 213.75:
		dir = "SSW"
	case deg >= 213.75 && deg < 236.25:
		dir = "SW"
	case deg >= 236.25 && deg < 258.75:
		dir = "WSW"
	case deg >= 258.75 && deg < 281.25:
		dir = "W"
	case deg >= 281.25 && deg < 303.75:
		dir = "WNW"
	case deg >= 303.75 && deg < 326.25:
		dir = "NW"
	case deg >= 326.25 && deg < 348.75:
		dir = "NNW"
	default:
		dir = "N"
	}
	dis *= 10
	dis = math.Floor(dis + 0.5)
	dis /= 10
	switch {
	case dis >= 0.0 && dis <= 0.2:
		w = 0
	case dis >= 0.3 && dis <= 1.5:
		w = 1
	case dis >= 1.6 && dis <= 3.3:
		w = 2
	case dis >= 3.4 && dis <= 5.4:
		w = 3
	case dis >= 5.5 && dis <= 7.9:
		w = 4
	case dis >= 8.0 && dis <= 10.7:
		w = 5
	case dis >= 10.8 && dis <= 13.8:
		w = 6
	case dis >= 13.9 && dis <= 17.1:
		w = 7
	case dis >= 17.2 && dis <= 20.7:
		w = 8
	case dis >= 20.8 && dis <= 24.4:
		w = 9
	case dis >= 24.5 && dis <= 28.4:
		w = 10
	case dis >= 28.5 && dis <= 32.6:
		w = 11
	case dis >= 32.7:
		w = 12
	}
	if w == 0 {
		dir = "C"
	}
	fmt.Println(dir, w)
}
