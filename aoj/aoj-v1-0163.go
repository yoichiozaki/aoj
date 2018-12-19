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
	distance_table := [][]int{
		{0, 6, 13, 18, 23, 43, 58},
		{6, 0, 7, 12, 17, 37, 52},
		{13, 7, 0, 5, 10, 30, 45},
		{18, 12, 5, 0, 5, 25, 40},
		{23, 17, 10, 5, 0, 20, 35},
		{43, 37, 30, 25, 20, 0, 15},
		{58, 52, 45, 40, 35, 15, 0},
	}
	toll_table := [][]int{
		{0, 300, 500, 600, 700, 1350, 1650},
		{300, 0, 350, 450, 600, 1150, 1500},
		{500, 350, 0, 250, 400, 1000, 1350},
		{600, 450, 250, 0, 250, 850, 1300},
		{700, 600, 400, 250, 0, 600, 1150},
		{1350, 1150, 1000, 850, 600, 0, 500},
		{1650, 1500, 1350, 1300, 1150, 500, 0},
	}
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		if d == 0 {
			break
		}
		d--
		scanner.Scan()
		dbuf := strings.Split(scanner.Text(), " ")
		dh, _ := strconv.Atoi(dbuf[0])
		dm, _ := strconv.Atoi(dbuf[1])
		dtime := 100*dh + dm
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		a--
		scanner.Scan()
		abuf := strings.Split(scanner.Text(), " ")
		ah, _ := strconv.Atoi(abuf[0])
		am, _ := strconv.Atoi(abuf[1])
		atime := 100*ah + am
		toll := toll_table[d][a]
		if distance_table[d][a] <= 40 &&
			1730 <= dtime && dtime <= 1930 ||
			1730 <= atime && atime <= 1930 {
			toll = ((toll/2-1)/50 + 1) * 50
		}
		fmt.Println(toll)
	}
}
