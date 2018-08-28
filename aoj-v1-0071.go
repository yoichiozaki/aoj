
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
var table [][]string

func main() {
	table = make([][]string, 8)
	for i := range table {
		table[i] = make([]string, 9)
	}
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	scanner.Text()
	now := 1
	for i := 0; i < n; i++ {
		for j := 0; j < 8; j++ {
			scanner.Scan()
			table[j] = strings.Split(scanner.Text(), "")
		}
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		// log.Println(table)
		// log.Printf("x = %d, y = %d\n", x, y)
		scanner.Scan()
		scanner.Text()
		solver(y-1, x-1)
		fmt.Printf("Data %d:\n", now)
		now++
		for i := 0; i < 8; i++ {
			fmt.Printf("%s\n", strings.Join(table[i], ""))
		}
	}
}

func solver(a, b int) {
	table[a][b] = "0"
	for i := 1; i <= 3; i++ {
		if a-i >= 0 && table[a-i][b] == "1" {
			solver(a-i, b)
		}
		if b-i >= 0 && table[a][b-i] == "1" {
			solver(a, b-i)
		}
		if a+i < 8 && table[a+i][b] == "1" {
			solver(a+i, b)
		}
		if b+i < 8 && table[a][b+i] == "1" {
			solver(a, b+i)
		}
	}
}
