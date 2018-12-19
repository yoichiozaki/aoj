
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	a, b, c := false, false, false
	score := 0
	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < d; i++ {
		score = 0
		a, b, c = false, false, false
		out := 0
		for {
			scanner.Scan()
			event := scanner.Text()
			if event == "HIT" {
				if c {
					score++
				}
				c = b
				b = a
				a = true
			} else if event == "HOMERUN" {
				score++
				if a {
					score++
				}
				if b {
					score++
				}
				if c {
					score++
				}
				a, b, c = false, false, false
			} else {
				out++
			}
			if out == 3 {
				fmt.Println(score)
				break
			}
		}
	}
}
