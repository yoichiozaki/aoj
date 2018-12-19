
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func rangeCheck(x, y int) bool {
	if x >= 0 && x <= 9 && y >= 0 && y <= 9 {
		return true
	} else {
		return false
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	paper := make([][]int, 10)
	for i := range paper {
		paper[i] = make([]int, 10)
		for j := range paper[i] {
			paper[i][j] = 0
		}
	}
	// fmt.Println(paper)
	for scanner.Scan() {
		infos := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(infos[0])
		y, _ := strconv.Atoi(infos[1])
		s, _ := strconv.Atoi(infos[2])
		// fmt.Println(x, y, s)
		switch s {
		case 1:
			paper[x][y]++
			if rangeCheck(x, y-1) {
				paper[x][y-1]++
			}
			if rangeCheck(x, y+1) {
				paper[x][y+1]++
			}
			if rangeCheck(x-1, y) {
				paper[x-1][y]++
			}
			if rangeCheck(x+1, y) {
				paper[x+1][y]++
			}
			continue
		case 2:
			paper[x][y]++
			if rangeCheck(x-1, y-1) {
				paper[x-1][y-1]++
			}
			if rangeCheck(x, y-1) {
				paper[x][y-1]++
			}
			if rangeCheck(x+1, y-1) {
				paper[x+1][y-1]++
			}
			if rangeCheck(x-1, y) {
				paper[x-1][y]++
			}
			if rangeCheck(x+1, y) {
				paper[x+1][y]++
			}
			if rangeCheck(x-1, y+1) {
				paper[x-1][y+1]++
			}
			if rangeCheck(x, y+1) {
				paper[x][y+1]++
			}
			if rangeCheck(x+1, y+1) {
				paper[x+1][y+1]++
			}
			continue
		case 3:
			paper[x][y]++
			if rangeCheck(x-1,y-1) {
				paper[x-1][y-1]++
			}
			if rangeCheck(x,y-1) {
				paper[x][y-1]++
			}
			if rangeCheck(x+1,y-1) {
				paper[x+1][y-1]++
			}
			if rangeCheck(x-1,y) {
				paper[x-1][y]++
			}
			if rangeCheck(x+1,y) {
				paper[x+1][y]++
			}
			if rangeCheck(x-1,y+1) {
				paper[x-1][y+1]++
			}
			if rangeCheck(x,y+1) {
				paper[x][y+1]++
			}
			if rangeCheck(x+1,y+1) {
				paper[x+1][y+1]++
			}
			if rangeCheck(x,y-2) {
				paper[x][y-2]++
			}
			if rangeCheck(x-2,y) {
				paper[x-2][y]++
			}
			if rangeCheck(x+2,y) {
				paper[x+2][y]++
			}
			if rangeCheck(x,y+2) {
				paper[x][y+2]++
			}
			continue
		}
	}
	count := 0
	max := -1
	for i := range paper {
		for j := range paper[i] {
			if paper[i][j] == 0 {
				count++
			}
			if max >= paper[i][j] {
				max = max
			} else {
				max = paper[i][j]
			}
		}
	}
	fmt.Printf("%d\n%d\n", count, max)
}
