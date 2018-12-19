
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		if scanner.Text() == "0" {
			break
		}
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		goal, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		d, _ := strconv.Atoi(scanner.Text())
		buf := make([]string, 0)
		for i := 0; i < d; i++ {
			scanner.Scan()
			buf = append(buf, scanner.Text())
		}
		// log.Println(buf)
		// log.Printf("n = %d, m = %d, goal = %d, d = %d", n, m, goal, d)
		info := make([][]string, len(buf))
		for i := 0; i < len(buf); i++ {
			info[i] = strings.Split(buf[i], "")
		}
		// log.Println(info)
		m--
		goal--
		flag := false
		if solver(m, n, d, info) == goal {
			fmt.Println(0)
			flag = true
		}
		for i := 0; !flag && i < d; i++ {
			for j := 0; !flag && j < n-1; j++ {
				if info[i][j] == "1" {
					continue
				}
				info[i][j] = "1"
				if solver(m, n, d, info) == goal {
					fmt.Printf("%d %d\n", i+1, j+1)
					flag = true
				}
				info[i][j] = "0"
			}
		}
		if !flag {
			fmt.Println(1)
		}
	}
}

func solver(t int, n int, d int, info [][]string) int {
	for i := 0; i < d; i++ {
		for j := 0; j < n-2; j++ {
			if info[i][j] == "1" && info[i][j+1] == "1" {
				return -1
			}
		}
	}
	for i := 0; i < d; i++ {
		if t < n-1 && info[i][t] == "1" {
			t++
		} else if t != 0 && info[i][t-1] == "1" {
			t--
		}
	}
	return t
}
