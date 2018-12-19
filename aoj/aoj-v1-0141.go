
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	out := make([][]string, 100)
	for i := range out {
		out[i] = make([]string, 101)
	}
	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	for a > 0 {
		a--
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < b; i++ {
			for j := 0; j < b; j++ {
				out[i][j] = " "
			}
			out[i][b] = "0"
		}
		if b == 1 {
			fmt.Println("#")
			if a != 0 {
				fmt.Println()
			}
			continue
		}
		if b == 2 {
			fmt.Printf("##\n# \n")
			if a != 0 {
				fmt.Println()
				continue
			}
		}
		row := b -1
		col := 0
		dir := 0
		l := 0
		out[row][col] = "#"
		for {
			nrow := row+dy[dir%4]
			ncol := col+dx[dir%4]
			if dir%4 == 0 && nrow == 0 || dir%4==3&&ncol==0 || dir%4==2&&nrow==b-1 || dir%4==1&&ncol==b-1 {
				row=nrow
				col=ncol
				out[row][col]="#"
				dir++
				l = 1
			} else if out[nrow+dy[dir%4]][ncol+dx[dir%4]]=="#" {
				if l < 3 {
					break
				}
				dir++
				l = 1
			} else {
				row=nrow
				col=ncol
				out[row][col]="#"
				l++
			}
		}
		for i := 0; i < b; i++ {
			for j := 0; j < b; j++ {
				fmt.Printf("%s", out[i][j])
			}
			fmt.Println()
		}
		if a != 0 {
			fmt.Println()
		}
	}
}
