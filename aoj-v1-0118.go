
import (
	"bufio"
		"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
var c = make([][]string, 100)
var h, w int

func main() {
	for {
		for i := range c {
			c[i] = make([]string, 100)
		}
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		h, _ = strconv.Atoi(input[0])
		w, _ = strconv.Atoi(input[1])
		if h == 0 && w == 0 {
			break
		}
		for i := 0; i < h; i++ {
			scanner.Scan()
			c[i] = strings.Split(scanner.Text(), "")
		}
		ret := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if c[i][j] != "." {
					ret++
					del(i, j, c[i][j])
				}
			}
		}
		fmt.Println(ret)
	}
}

func del(a, b int, t string) {
	c[a][b] = "."
	if a>0&&c[a-1][b]==t {
		del(a-1,b,t)
	}
	if b>0&&c[a][b-1]==t {
		del(a,b-1,t)
	}
	if a<h-1&&c[a+1][b]==t {
		del(a+1,b,t)
	}
	if b<w-1&&c[a][b+1]==t {
		del(a,b+1,t)
	}
}

