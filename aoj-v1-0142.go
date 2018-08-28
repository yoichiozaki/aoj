
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		table := make([]bool, 10001)
		now := make([]int, 10001)
		count := make([]int, 10001)
		a, _ := strconv.Atoi(scanner.Text())
		if a == 0 {
			break
		}
		index := 0
		for i := 1; i < a; i++ {
			val := i*i%a
			if !table[val] {
				table[val] = true
				now[index] = val
				index++
			}
		}
		for i := 0; i < index; i++ {
			for j := i+1; j < index; j++ {
				t := now[i] - now[j]
				if t < 0 {
					t += a
				}
				if t > (a-1)/2 {
					t = a - t
				}
				count[t]++
			}
		}
		for i := 1; i <= a/2; i++ {
			fmt.Println(count[i]*2)
		}
	}
}
