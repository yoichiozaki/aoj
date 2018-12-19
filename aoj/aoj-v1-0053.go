
import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

var table [500000]int
var prime [100000]int
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for i := 0; i < 500000; i++ {
		table[i] = 0
	}
	for i := 0; i < 100000; i++ {
		prime[i] = 0
	}
	table[0] = -1
	table[1] = -1
	for i := 2; i < 500000; i++ {
		if table[i] != -1 {
			table[i] = 1
			for j := 2*i; j < 500000; j += i {
				table[j] = -1
			}
		}
	}
	now := 0
	for i := 0; i < 500000; i++ {
		if table[i] == 1 {
			prime[now] = i
			now++
		}
	}
	for scanner.Scan() {
		if scanner.Text() == "0" {
			break
		}
		a, _ := strconv.Atoi(scanner.Text())
		ret := 0
		for i := 0; i < a; i++ {
			ret += prime[i]
		}
		fmt.Printf("%d\n", ret)
	}
}
