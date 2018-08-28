
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

/*********************************************************************
長さmの列X_(m)と長さnの列Y_(n)の最長共通部分列LCSは、
	(1)X[m] == Y[n]ならば、X_(m-1)とY_(n-1)のLCSにX[m]を繋げたもの
	(2)X[m] != Y[n]ならば、X_(m-1)とY_(n)のLCSもしくはX_(m)とY_(n-1)のLCS
	のいずれかのうち長い方

c[m+1][n+1]：c[i][j]にX_(i)とY_(j)のLCSの長さを格納する二次元配列
*********************************************************************/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		s1 := scanner.Text()
		scanner.Scan()
		s2 := scanner.Text()
		s1, s2 = " " + s1 , " " + s2
		c := make([][]int, len(s1))
		for i := range c {
			c[i] = make([]int, len(s2))
		}

		LCS := 0

		for i := 1; i < len(s1); i++ {
			for j := 1; j < len(s2); j++ {
				if s1[i] == s2[j] {
					c[i][j] = c[i-1][j-1] + 1
				} else {
					c[i][j] = max(c[i-1][j], c[i][j-1])
				}
				LCS = max(LCS, c[i][j])
			}
		}
		fmt.Println(LCS)
	}
}
