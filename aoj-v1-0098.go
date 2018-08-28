
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	a := make([][]int, 102)
	for i := range a {
		a[i] = make([]int, 102)
	}
	s := make([][]int, 102)
	for i := range s {
		s[i] = make([]int, 102)
	}
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmp := strings.Split(scanner.Text(), " ")
		for j := range tmp {
			a[i][j], _ = strconv.Atoi(tmp[j])
		}
	}
	// log.Println(a)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			s[r][c+1] += s[r][c]+a[r][c]
		}
	}
	ans := s[0][1]
	for c := 0; c < n; c++ {
		for k := c+1; k < n+1; k++ {
			t := 0
			for r := 0; r < n; r++ {
				if t < 0 {
					t = s[r][k]-s[r][c]
				} else {
					t += s[r][k]-s[r][c]
				}
				if t > ans {
					ans = t
				}
			}
		}
	}
	fmt.Println(ans)
}
