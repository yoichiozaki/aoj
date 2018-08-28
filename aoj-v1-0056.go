
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var table [500001]int
	for i := 0; i < 500001; i++ {
		table[i] = 0
	}
	table[1] = -1
	table[0] = -1
	for i := 2; i < 500001; i++ {
		if table[i] == 0 {
			table[i] = 1
			for j := i*2; j < 500001; j += i {
				table[j] = -1
			}
		}
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		ret := 0
		for i := 1; i <= n/2; i++ {
			if table[i] == 1 && table[n-i] == 1 {
				ret++
			}
		}
		fmt.Println(ret)
	}
}
