// Great-circular distance
// Reference: https://ja.wikipedia.org/wiki/大円距離

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

const (
	R = 6378.1
	M = math.Pi / 180
)

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		input := make([]float64, len(buf))
		for i := range input {
			input[i], _ = strconv.ParseFloat(buf[i], 64)
		}
		if input[0] == -1 && input[1] == -1 &&
			input[2] == -1 && input[3] == -1 {
			break
		}
		y1, x1, y2, x2 := input[0], input[1], input[2], input[3]
		a := math.Cos(y1*M)*math.Cos(y2*M)*math.Cos((x1-x2)*M) +
			math.Sin(y1*M)*math.Sin(y2*M)
		fmt.Println(int(R*math.Acos(a) + 0.5))
	}
}
