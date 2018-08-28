
import (
	"bufio"
	"os"
	"fmt"
	"strings"
	)

var scanner = bufio.NewScanner(os.Stdin)
var title = [4]int{0, 90, 180, 270}

func main() {
	p := make([][]string, 8)
	for i := range p {
		scanner.Scan()
		p[i] = strings.Split(scanner.Text(), "")
	}
	// log.Println(p)
	for i := 1; i < 4; i++ {
		fmt.Println(title[i])
		p = spin(p)
		print_p(p)
	}
}

func spin(p [][]string) [][]string {
	ret := make([][]string, 8)
	for i := range ret {
		ret[i] = make([]string, 8)
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			ret[i][j] = p[7-j][i]
		}
	}
	return ret
}

func print_p(p [][]string) {
	for _, v := range p {
		fmt.Println(strings.Join(v, ""))
	}
}
