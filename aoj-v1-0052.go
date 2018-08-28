
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		if scanner.Text() == "0" {
			break
		}
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(get0(n))
	}
}

func get0(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i % 5 == 0 {
			count++
		}
	}
	for i := 1; i <= n; i++ {
		if i % 25 == 0 {
			count++
		}
	}
	for i := 1; i <= n; i++ {
		if i % 125 == 0 {
			count++
		}
	}
	for i := 1; i <= n; i++ {
		if i % 625 == 0 {
			count++
		}
	}
	for i := 1; i <= n; i++ {
		if i % 3125 == 0 {
			count++
		}
	}
	for i := 1; i <= n; i++ {
		if i % 15625 == 0 {
			count++
		}
	}
	return count
}

