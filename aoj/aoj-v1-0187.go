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

type Point struct {
	x, y float64
}

func main() {
	for scanner.Scan() {
		buf1 := strings.Split(scanner.Text(), " ")
		line1 := make([]float64, len(buf1))
		for i := range line1 {
			line1[i], _ = strconv.ParseFloat(buf1[i], 64)
		}
		if line1[0] == 0 && line1[1] == 0 &&
			line1[2] == 0 && line1[3] == 0 {
			break
		}
		scanner.Scan()
		buf2 := strings.Split(scanner.Text(), " ")
		line2 := make([]float64, len(buf2))
		for i := range line2 {
			line2[i], _ = strconv.ParseFloat(buf2[i], 64)
		}
		scanner.Scan()
		buf3 := strings.Split(scanner.Text(), " ")
		line3 := make([]float64, len(buf3))
		for i := range line2 {
			line3[i], _ = strconv.ParseFloat(buf3[i], 64)
		}
		if !intersect(line1[0], line1[1], line1[2], line1[3], line2[0], line2[1], line2[2], line2[3]) ||
			!intersect(line2[0], line2[1], line2[2], line2[3], line3[0], line3[1], line3[2], line3[3]) ||
			!intersect(line3[0], line3[1], line3[2], line3[3], line1[0], line1[1], line1[2], line1[3]) {
			fmt.Println("kyo")
		} else {
			p1 := intersection(line1[0], line1[1], line1[2], line1[3], line2[0], line2[1], line2[2], line2[3])
			p2 := intersection(line2[0], line2[1], line2[2], line2[3], line3[0], line3[1], line3[2], line3[3])
			p3 := intersection(line3[0], line3[1], line3[2], line3[3], line1[0], line1[1], line1[2], line1[3])
			area := math.Abs((p2.x-p1.x)*(p3.y-p1.y)-(p2.y-p1.y)*(p3.x-p1.x)) / 2
			if area < 0.00000001 {
				fmt.Println("kyo")
			} else if 1899999.99999999 < area {
				fmt.Println("dai-kichi")
			} else if 999999.99999999 < area {
				fmt.Println("chu-kichi")
			} else if 99999.99999999 < area {
				fmt.Println("kichi")
			} else {
				fmt.Println("syo-kichi")
			}
		}
	}
}

func cross(x1, y1, x2, y2 float64) int {
	z := x1*y2 - y1*x2
	if 0 < z {
		return 1
	} else if z < 0 {
		return -1
	} else {
		return 0
	}
}

func intersect(x1, y1, x2, y2, x3, y3, x4, y4 float64) bool {
	if 0 <= cross(x2-x1, y2-y1, x3-x1, y3-y1)*
		cross(x2-x1, y2-y1, x4-x1, y4-y1) {
		return false
	}
	if 0 <= cross(x4-x3, y4-y3, x1-x3, y1-y3)*
		cross(x4-x3, y4-y3, x2-x3, y2-y3) {
		return false
	}
	return true
}

func intersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) Point {
	if x3 == x4 {
		x1, x3 = x3, x1
		y1, y3 = y3, y1
		x2, x4 = x4, x2
		y2, y4 = y4, y2
	}
	ret := Point{}
	if x1 == x2 {
		// y軸に平行な直線の場合
		c := (y4 - y3) / (x4 - x3)
		d := y3 - c*x3
		x := x1
		y := c*x + d
		ret.x, ret.y = x, y
	} else {
		a := (y2 - y1) / (x2 - x1)
		b := y1 - a*x1
		c := (y4 - y3) / (x4 - x3)
		d := y3 - c*x3
		x := (d - b) / (a - c)
		y := a*x + b
		ret.x, ret.y = x, y
	}
	return ret
}
