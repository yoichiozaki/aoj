package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		m, _ := strconv.Atoi(scanner.Text())
		if m == 0 {
			break
		}
		args1 := make([]float64, m)
		x := 0.0
		for i := 0; i < m-1; i++ {
			scanner.Scan()
			tmp, _ := strconv.ParseFloat(scanner.Text(), 64)
			x += tmp
			args1[i] = tmp
		}
		args1[m-1] = 360 - x
		first_object := 0.0
		for _, v := range args1 {
			first_object += math.Sin(v * math.Pi / 180)
		}
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		args2 := make([]float64, n)
		y := 0.0
		for i := 0; i < n-1; i++ {
			scanner.Scan()
			tmp, _ := strconv.ParseFloat(scanner.Text(), 64)
			y += tmp
			args2[i] = tmp
		}
		args2[n-1] = 360 - y
		second_object := 0.0
		for _, v := range args2 {
			second_object += math.Sin(v * math.Pi / 180)
		}
		if first_object > second_object+0.0000001 {
			fmt.Println(1)
		} else if second_object > first_object+0.0000001 {
			fmt.Println(2)
		} else {
			fmt.Println(0)
		}
	}
}
