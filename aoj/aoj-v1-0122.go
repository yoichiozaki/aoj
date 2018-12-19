
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	move := [][]int{
		{-1, -2},
		{0, -2},
		{1, -2},
		{2, -1},
		{2, 0},
		{2, 1},
		{1, 2},
		{0, 2},
		{-1, 2},
		{-2, 1},
		{-2, 0},
		{-2, -1},
	}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		px, _ := strconv.Atoi(input[0])
		py, _ := strconv.Atoi(input[1])
		if px == 0 && py == 0 {
			break
		}
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		xy := strings.Split(scanner.Text(), " ")
		arr := make([][][]int, 10)
		for i := range arr {
			arr[i] = make([][]int, 10)
			for j := range arr[i] {
				arr[i][j] = make([]int, 10)
			}
		}
		for i := 0; i < n; i++ {
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			xy = xy[2:]
			for nx := x-1; nx < x+2; nx++ {
				for ny := y-1; ny < y+2; ny++ {
					if 0 <= nx && 0 <= ny && nx <= 9 && ny <= 9 {
						arr[i][nx][ny] = 1
					}
				}
			}
		}
		q := make([]entry, 0)
		q = append(q, entry{px, py, 0})
		flag := true
		for flag && len(q) != 0 {
			tmp := q[0]
			q = q[1:]
			x := tmp.px
			y := tmp.py
			k := tmp.k
			for i := 0; i < 12; i++ {
				nx, ny := x + move[i][0], y + move[i][1]
				if 0 <= nx && 0 <= ny && nx <= 9 && ny <= 9 && arr[k][nx][ny] != 0 {
					if k+1 < n {
						q = append(q, entry{nx, ny, k+1})
					} else {
						fmt.Println("OK")
						flag = false
						break
					}
				}
			}
		}
		if flag {
			fmt.Println("NA")
		}
	}
}

type entry struct {
	px, py, k int
}

