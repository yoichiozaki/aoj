
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

const NUM = 60000

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var table [NUM]int
	var limit int
	for i := 0; i < NUM; i++ {
		table[i] = 1
	}
	table[0] = 0
	limit = int(math.Sqrt(NUM))
	for i := 2; i <= limit; i++ {
		if table[i] == 1 {
			for k := 2*i; k < NUM; k += i {
				table[k] = 0
			}
		}
	}

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		for k := n-1; k >= 0; k-- {
			if table[k] == 1 {
				fmt.Printf("%d ", k)
				break
			}
		}
		for k := n+1; k < NUM; k++ {
			if table[k] == 1 {
				fmt.Printf("%d\n", k)
				break
			}
		}
	}
}
