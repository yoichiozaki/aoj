
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
		"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	a := make([]int, 12)
	b := make([]int, 12)
	d := make([]int, 12)
	p := make([]int, 30)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		for i := 0; i < n; i++ {
			a[i], _ = strconv.Atoi(input[i])
			b[i] = a[i]
		}
		count := 0
		for {
			for i := 0; i < 30; i++ {
				p[i] = 0
			}
			for i := 0; i < n; i++ {
				p[b[i]]++
			}
			for i := 0; i < n; i++ {
				d[i] = p[b[i]]
			}
			ok := true
			for i := 0; i < n; i++ {
				if d[i] - b[i] != 0 {
					ok = false
				}
			}
			if ok {
				break
			}
			for i := 0; i < n; i++ {
				b[i] = d[i]
			}
			count++
		}
		fmt.Println(count)
		for i := 0; i < n; i++ {
			if i != 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", d[i])
		}
		fmt.Println()
	}
}

