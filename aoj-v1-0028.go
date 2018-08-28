
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	table := make([]int, 100)
	now := 0
	max := 0
	for i := range table {
		table[i] = 0
	}
	for scanner.Scan() {
		now, _ =strconv.Atoi(scanner.Text())
		table[now-1]++
		if table[now-1] > max {
			max = table[now-1]
		}
	}
	for i := 0; i < 100; i++ {
		if table[i] == max {
			fmt.Println(i+1)
		}
	}
}
