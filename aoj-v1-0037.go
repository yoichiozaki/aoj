
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var points [5][5][4]bool

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for i := 0; i < 9; i++ {
		y := i/2
		if i%2 == 0 {
			scanner.Scan()
			info := strings.Split(scanner.Text(), "")
			for k := range info {
				e, _ := strconv.Atoi(info[k])
				if e != 0 {
					points[y][k][1], points[y][k+1][3] = true, true
				}
			}
		} else {
			scanner.Scan()
			info := strings.Split(scanner.Text(), "")
			for k := range info {
				e, _ := strconv.Atoi(info[k])
				if e != 0 {
					points[y][k][2], points[y+1][k][0] = true, true
				}
			}
		}
	}
	x := 1
	y := 0
	dir := 1
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	ds := "URDL"
	fmt.Printf("%s", string(ds[dir]))
	for !(x == 0 && y == 0) {
		if points[y][x][(dir+3)%4] {
			dir = (dir+3)%4
		} else if points[y][x][dir] {

		} else if points[y][x][(dir+1)%4] {
			dir = (dir+1)%4
		} else {
			dir = (dir+2)%4
		}
		x += dx[dir]
		y += dy[dir]
		fmt.Printf("%s", string(ds[dir]))
	}
	fmt.Println()
}
