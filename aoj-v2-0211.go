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
		d, v := make([]int, 0), make([]int, 0)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(buf[0])
			b, _ := strconv.Atoi(buf[1])
			g := _gcd(a, b)
			a /= g
			b /= g
			d = append(d, a)
			v = append(v, b)
		}
		g := lcm(n, d)
		for i := 0; i < n; i++ {
			d[i] = (g / d[i]) * v[i]
		}
		g = gcd(n, d)
		for i := 0; i < n; i++ {
			fmt.Println(d[i] / g)
		}
	}
}

func _gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a, b = b, r
	}
	return a
}

func gcd(n int, a []int) int {
	if n == 1 {
		return a[0]
	}
	g := _gcd(a[0], a[1])
	for i := 2; i < n; i++ {
		if g == 1 {
			break
		}
		g = _gcd(g, a[i])
	}
	return g
}

func lcm(n int, a []int) int {
	if n == 1 {
		return a[0]
	}
	g := _gcd(a[0], a[1])
	c := a[0] / g * a[1]
	for i := 2; i < n; i++ {
		g = _gcd(c, a[i])
		c = c / g * a[i]
	}
	return c
}
