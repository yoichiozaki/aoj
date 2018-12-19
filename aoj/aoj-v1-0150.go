
import (
	"bufio"
	"os"
		"strconv"
	"fmt"
	)

var scanner= bufio.NewScanner(os.Stdin)

func main() {
	MAX := 10001
	SQRT := 100
	isPrime := make([]bool, MAX)
	for i :=range isPrime {
		isPrime[i] = true
	}
	for i := 3; i < SQRT; i += 2 {
		if isPrime[i] {
			for j := i*i; j < MAX; j += i {
				isPrime[j] = false
			}
		}
	}
	// log.Println(isPrime)
	tbl := make([]int, MAX)
	ans := 0
	for i := 5; i < MAX; i += 2 {
		if isPrime[i] && isPrime[i-2] {
			ans = i
		}
		tbl[i] = ans
	}
	// log.Println(tbl)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		if n&1 == 0 {
			n--
		}
		fmt.Printf("%d %d\n", tbl[n]-2, tbl[n])
	}
}
