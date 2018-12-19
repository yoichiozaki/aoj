
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	t := make([][]int, 1000)
	for i := range t {
		t[i] = make([]int, 1000)
	}
	table := make([][]int, 1000)
	for i := range table {
		table[i] = make([]int, 1000)
	}
	L := make([][]int, 1000)
	for i := range L {
		L[i] = make([]int, 1000)
	}
	R := make([][]int, 1000)
	for i := range R {
		R[i] = make([]int, 1000)
	}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		h, _ := strconv.Atoi(input[0])
		w, _ := strconv.Atoi(input[1])
		if h == 0 {
			break
		}
		for i := 0; i < h; i++ {
			scanner.Scan()
			y := strings.Split(scanner.Text(), "")
			for j := 0; j < w; j++ {
				if y[j] == "." {
					t[i][j] = 1
				} else {
					t[i][j] = 0
				}
			}
		}
		for i := 0; i < w; i++ {
			now := 0
			for j := 0; j < h; j++ {
				if t[j][i] == 1 {
					now++
				} else {
					now = 0
				}
				table[j][i] = now
			}
		}
		for i := 0; i < h; i++ {
			stack := make([]int, 0)
			for j := 0; j < w; j++ {
				for len(stack) != 0 && table[i][stack[len(stack)-1]] >= table[i][j] {
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					L[i][j] = 0
				} else {
					L[i][j] = stack[len(stack)-1] + 1
				}
				stack = append(stack, j)
			}
			for len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			for j := w-1; j >= 0; j-- {
				for len(stack) != 0 && table[i][stack[len(stack)-1]] >= table[i][j] {
					stack = stack[:len(stack)-1]
				}
				if len(stack) == 0 {
					R[i][j] = w-1
				} else {
					R[i][j] = stack[len(stack)-1] - 1
				}
				stack = append(stack, j)
			}
		}
		ret := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				ret = max(ret, table[i][j]*(R[i][j]-L[i][j]+1))
			}
		}
		fmt.Println(ret)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
